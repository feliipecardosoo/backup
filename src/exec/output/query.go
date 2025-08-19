package output

import (
	"backup-etl/src/config/model"
	"context"
)

// Querier define a interface pública que será exposta pelo pacote output.
// Qualquer implementação de Queries deve fornecer os métodos abaixo.
type Querier interface {
	// insertUserToday insere os usuários do dia no banco de backup
	insertUserToday(ctx context.Context) error

	// DeleteExistingUsers remove usuários que já existem no banco de backup
	DeleteExistingUsers(ctx context.Context, users []model.Membro) (int64, error)
}
