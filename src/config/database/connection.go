package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoConnection define a interface para gerenciamento da conexão com o MongoDB.
// Ela abstrai as operações básicas necessárias para conectar ao banco e acessar coleções.
type MongoConnection interface {
	// Connect estabelece uma conexão com o MongoDB usando a URI fornecida.
	// Retorna erro caso a conexão falhe.
	Connect(uri string) error

	// Collection retorna uma referência para a coleção especificada no banco especificado.
	// Permite realizar operações de leitura e escrita nesta coleção.
	Collection(dbName, collectionName string) *mongo.Collection

	// ContextWithTimeout cria e retorna um contexto com timeout predefinido (ex: 15 segundos),
	// útil para limitar o tempo de execução de operações que acessam o banco.
	ContextWithTimeout() (context.Context, context.CancelFunc)

	// Disconnect encerra a conexão com o MongoDB utilizando o contexto informado.
	// Retorna erro caso ocorra falha durante a desconexão.
	Disconnect(ctx context.Context) error
}
