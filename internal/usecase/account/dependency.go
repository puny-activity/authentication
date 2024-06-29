package account

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/role"
)

type accountRepository interface {
	IsUsernameTakenTx(ctx context.Context, tx *sqlx.Tx, username string) (bool, error)
	CreateTx(ctx context.Context, tx *sqlx.Tx, accountToCreate account.ToCreateWithHashedPassword) error
	AccountsCountTx(ctx context.Context, tx *sqlx.Tx) (int, error)
}

type roleRepository interface {
	AssignTx(ctx context.Context, tx *sqlx.Tx, accountID uuid.UUID, role role.Role) error
}
