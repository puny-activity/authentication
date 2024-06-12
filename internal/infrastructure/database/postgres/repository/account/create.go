package account

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity"
	"github.com/puny-activity/authentication/internal/interr"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) Create(ctx context.Context, accountToCreate entity.AccountCreateRequestWithHashedPassword) error {
	return r.create(ctx, r.db, accountToCreate)
}

func (r Repository) CreateTx(ctx context.Context, tx *sqlx.Tx, accountToCreate entity.AccountCreateRequestWithHashedPassword) error {
	return r.create(ctx, tx, accountToCreate)
}

func (r Repository) create(ctx context.Context, queryer queryer.Queryer, accountToCreate entity.AccountCreateRequestWithHashedPassword) error {
	query := `
INSERT INTO accounts(id, username, hashed_password, created_at, last_online)
VALUES ($1, $2, $3, $4, $5)
`

	_, err := queryer.ExecContext(ctx, query,
		accountToCreate.ID,
		accountToCreate.Username,
		accountToCreate.HashedPassword,
		accountToCreate.CreatedAt.ToDateTimeString(),
		nil)
	if err != nil {
		return werr.WrapEE(interr.DatabaseFailedToExecuteQuery, err)
	}

	return nil
}
