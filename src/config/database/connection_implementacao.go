package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoConnectionImpl é a implementação concreta da interface MongoConnection,
// responsável por gerenciar a conexão com o banco MongoDB.
type mongoConnectionImpl struct {
	client *mongo.Client
}

// NewMongoConnection cria uma nova instância da implementação de MongoConnection.
func NewMongoConnection() MongoConnection {
	return &mongoConnectionImpl{}
}

// Connect estabelece conexão com o MongoDB utilizando a URI fornecida.
// Realiza o ping para garantir que a conexão foi estabelecida com sucesso.
// Retorna erro caso a conexão falhe.
func (m *mongoConnectionImpl) Connect(uri string) error {
	clientOptions := options.Client().ApplyURI(uri)

	// Cria um contexto com timeout de 10 segundos para limitar o tempo de conexão
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Tenta conectar ao MongoDB com as opções especificadas
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("erro ao conectar: %w", err)
	}

	// Testa a conexão com um ping
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("erro ao pingar o banco: %w", err)
	}

	fmt.Println("✅ Conectado ao MongoDB com sucesso!")
	m.client = client
	return nil
}

// Collection retorna uma referência para a coleção especificada no banco especificado.
// Permite realizar operações CRUD nesta coleção.
func (m *mongoConnectionImpl) Collection(dbName, collectionName string) *mongo.Collection {
	return m.client.Database(dbName).Collection(collectionName)
}

// ContextWithTimeout retorna um contexto com timeout de 15 segundos,
// usado para operações que precisam ser canceladas se demorarem muito.
func (m *mongoConnectionImpl) ContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}

// Disconnect encerra a conexão com o MongoDB utilizando o contexto fornecido.
// Retorna erro caso a desconexão falhe.
func (m *mongoConnectionImpl) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
