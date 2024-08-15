package account

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
)

type accountRepository interface {
	IsEmailTakenTx(ctx context.Context, tx *sqlx.Tx, email string) (bool, error)
	IsNicknameTakenTx(ctx context.Context, tx *sqlx.Tx, nickname string) (bool, error)
	CreateTx(ctx context.Context, tx *sqlx.Tx, accountToCreate account.ToCreateWithHashedPassword) error
	AccountsCountTx(ctx context.Context, tx *sqlx.Tx) (int, error)
}
