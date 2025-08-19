package output

import (
	"backup-etl/src/config/database"
	"backup-etl/src/config/model"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Queries gerencia as operações de inserção e exclusão no banco de backup.
// Ele recebe uma conexão MongoConnection já existente.
type Queries struct {
	conn database.MongoConnection // Conexão com o MongoDB
	db   string                   // Nome do banco de backup
	col  string                   // Nome da coleção de backup
}

// New cria uma nova instância de Queries usando uma conexão já existente.
// Lê as variáveis de ambiente MONGO_DB_BACKUP e MONGO_COLLECTION_BACKUP.
// Panica se as variáveis não estiverem configuradas.
func New(conn database.MongoConnection) *Queries {
	db := os.Getenv("MONGO_DB_BACKUP")
	col := os.Getenv("MONGO_COLLECTION_BACKUP")

	if db == "" || col == "" {
		panic("❌ Variáveis de ambiente MONGO_DB_BACKUP e/ou MONGO_COLLECTION_BACKUP não configuradas")
	}

	return &Queries{
		conn: conn,
		db:   db,
		col:  col,
	}
}

// collection retorna a referência da coleção de backup no MongoDB
func (q *Queries) collection() *mongo.Collection {
	return q.conn.Collection(q.db, q.col)
}

// InsertUsersToday insere um slice de usuários no banco de backup.
// Se o slice estiver vazio, apenas loga uma mensagem e retorna.
func (q *Queries) InsertUsersToday(ctx context.Context, users []model.Membro) error {
	if len(users) == 0 {
		log.Println("⚠️ Nenhum usuário para inserir no backup hoje.")
		return nil
	}

	coll := q.collection()

	// Converte slice de Membro para slice de interface{} exigido pelo InsertMany
	var docs []interface{}
	for _, u := range users {
		docs = append(docs, u)
	}

	_, err := coll.InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	log.Printf("✅ Inseridos %d usuários no backup.", len(users))
	return nil
}

// DeleteExistingUsers deleta do banco de backup todos os usuários cujo "Name"
// já existe no slice fornecido. Retorna a quantidade de documentos deletados.
func (q *Queries) DeleteExistingUsers(ctx context.Context, users []model.Membro) (int64, error) {
	if len(users) == 0 {
		return 0, nil
	}

	coll := q.collection()

	// Cria slice de nomes para filtro
	var names []string
	for _, u := range users {
		names = append(names, u.Name) // Ajustar se o campo for outro
	}

	// Executa exclusão em lote usando $in
	res, err := coll.DeleteMany(ctx, bson.M{"name": bson.M{"$in": names}})
	if err != nil {
		return 0, err
	}

	log.Printf("🗑️ Deletados %d usuários do backup.", res.DeletedCount)
	return res.DeletedCount, nil
}
