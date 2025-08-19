package input

import (
	"backup-etl/src/config/model"
	"context"
)

// Querier define a interface pública que será exposta pelo pacote input.
// Qualquer implementação deve fornecer um método para obter os usuários do dia.
type Querier interface {
	// GetUsersHoje retorna todos os usuários cuja data de modificação
	// seja igual à data de hoje.
	GetUsersHoje(ctx context.Context) ([]model.Membro, error)
}
