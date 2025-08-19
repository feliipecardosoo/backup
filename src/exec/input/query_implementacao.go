package input

import (
	"backup-etl/src/config/database"
	"backup-etl/src/config/model"
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Queries implementa a interface Querier, lidando com leitura
// do banco de input (banco inicial).
type Queries struct {
	conn database.MongoConnection // Conexão já existente com MongoDB
	db   string                  // Nome do banco de dados inicial
	col  string                  // Nome da coleção de membros
}

// New cria uma nova instância de Queries usando a conexão já existente.
// Lê as variáveis de ambiente MONGO_DB_NAME_INICIAL e MONGO_COLLECTION_MEMBRO_INICIAL.
// Panica se as variáveis não estiverem configuradas.
func New(conn database.MongoConnection) *Queries {
	db := os.Getenv("MONGO_DB_NAME_INICIAL")
	col := os.Getenv("MONGO_COLLECTION_MEMBRO_INICIAL")

	if db == "" || col == "" {
		panic("❌ Variáveis de ambiente MONGO_DB_NAME_INICIAL e/ou MONGO_COLLECTION_MEMBRO_INICIAL não configuradas")
	}

	return &Queries{
		conn: conn,
		db:   db,
		col:  col,
	}
}

// collection retorna a referência da coleção de membros no banco inicial
func (q *Queries) collection() *mongo.Collection {
	return q.conn.Collection(q.db, q.col)
}

// GetUsersHoje retorna todos os usuários cuja data de modificação
// seja igual à data de hoje, no formato dd/MM/yyyy.
// Utiliza o filtro "dataModificacao" na query.
func (q *Queries) GetUsersHoje(ctx context.Context) ([]model.Membro, error) {
	coll := q.collection()

	// Define explicitamente o fuso horário de São Paulo
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return nil, err
	}

	// Pega a data de hoje no formato dd/MM/yyyy no fuso horário correto
	hoje := time.Now().In(loc).Format("02/01/2006")
	log.Printf("Data usada no filtro: %s", hoje)

	// Busca usuários com dataModificacao igual à data de hoje
	cur, err := coll.Find(ctx, bson.M{
		"dataModificacao": hoje,
	})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []model.Membro
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
