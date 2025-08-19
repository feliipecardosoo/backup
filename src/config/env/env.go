package env

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv carrega as variáveis de ambiente definidas no arquivo .env
// localizado na raiz do projeto.
//
// Em caso de erro ao carregar o arquivo, a função encerra a aplicação
// com log fatal, pois as variáveis de ambiente são essenciais para o funcionamento.
//
// Essa função deve ser chamada no início da aplicação para garantir que
// todas as variáveis estejam disponíveis no ambiente de execução.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}
