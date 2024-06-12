package account

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity"
)

type accountRepository interface {
	IsUsernameTakenTx(ctx context.Context, tx *sqlx.Tx, username string) (bool, error)
	CreateTx(ctx context.Context, tx *sqlx.Tx, accountToCreate entity.AccountCreateRequestWithHashedPassword) error
	AccountsCountTx(ctx context.Context, tx *sqlx.Tx) (int, error)
}

type roleRepository interface {
	AssignTx(ctx context.Context, tx *sqlx.Tx, accountID string, role entity.Role) error
}
