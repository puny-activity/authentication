package account

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

type createDBParameter struct {
	ID             string  `db:"id"`
	Username       string  `db:"username"`
	Nickname       string  `db:"nickname"`
	HashedPassword string  `db:"hashed_password"`
	CreatedAt      string  `db:"created_at"`
	LastOnline     *string `db:"last_online"`
}

func (r Repository) Create(ctx context.Context, accountToCreate account.ToCreateWithHashedPassword) error {
	return r.create(ctx, r.db, accountToCreate)
}

func (r Repository) CreateTx(ctx context.Context, tx *sqlx.Tx, accountToCreate account.ToCreateWithHashedPassword) error {
	return r.create(ctx, tx, accountToCreate)
}

func (r Repository) create(ctx context.Context, queryer queryer.Queryer, accountToCreate account.ToCreateWithHashedPassword) error {
	query := `
INSERT INTO accounts(id, username, nickname, hashed_password, created_at, last_online)
VALUES (:id, :username, :nickname, :hashed_password, :created_at, :last_online)
`

	parameter := createDBParameter{
		ID:             accountToCreate.ID.String(),
		Username:       accountToCreate.Username,
		Nickname:       accountToCreate.Nickname,
		HashedPassword: accountToCreate.HashedPassword,
		CreatedAt:      accountToCreate.CreatedAt.ToDateTimeString(),
		LastOnline:     nil,
	}

	_, err := queryer.NamedExecContext(ctx, query, parameter)
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return nil
}
