package main

import (
	"backup-etl/src/config/database" // pacotes internos para conexão com MongoDB
	"backup-etl/src/exec/input"      // lógica para leitura do banco de input
	"backup-etl/src/exec/output"     // lógica para escrita no banco de backup
	"context"
	"log"
	"os"
	"time"
)

func main() {
	// Carrega variáveis do arquivo .env para o ambiente
	//env.LoadEnv()

	// Contexto global usado em todas as operações MongoDB
	ctx := context.Background()

	// -----------------------------
	// Conexão com o Mongo de input
	// -----------------------------
	// Lê a URI do MongoDB de input a partir das variáveis de ambiente
	bancoInicial := os.Getenv("BANCO_INICIAL")
	if bancoInicial == "" {
		log.Fatal("❌ BANCO_INICIAL não configurado")
	}

	// Cria e conecta a instância do MongoConnection para input
	connInput := database.NewMongoConnection()
	if err := connInput.Connect(bancoInicial); err != nil {
		log.Fatal(err)
	}
	// Garante que a conexão será encerrada ao final da execução
	defer connInput.Disconnect(ctx)

	// Cria o repositório de queries para input
	qInput := input.New(connInput)

	// -----------------------------
	// Pega os usuários do dia
	// -----------------------------
	users, err := qInput.GetUsersHoje(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("👀 Encontrados %d usuários no banco de input.", len(users))

	// Marca o início do processamento de backup
	start := time.Now()

	// -----------------------------
	// Conexão com o Mongo de output (backup)
	// -----------------------------
	connOutput := database.NewMongoConnection()
	if err := connOutput.Connect(os.Getenv("BANCO_BACKUP")); err != nil {
		log.Fatal(err)
	}
	defer connOutput.Disconnect(ctx)

	// Cria o repositório de queries para output
	qOutput := output.New(connOutput)

	// -----------------------------
	// Deleta usuários existentes no backup
	// -----------------------------
	// Garante que não haverá duplicados pelo campo "name" antes de inserir
	deleted, err := qOutput.DeleteExistingUsers(ctx, users)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("🗑️ %d usuários antigos removidos do backup.", deleted)

	// -----------------------------
	// Insere os usuários novos
	// -----------------------------
	if err := qOutput.InsertUsersToday(ctx, users); err != nil {
		log.Fatal(err)
	}

	// Calcula e exibe o tempo total gasto em todo o processo
	elapsed := time.Since(start)
	log.Printf("⏱️ Tempo total para executar o script: %s", elapsed)
}
