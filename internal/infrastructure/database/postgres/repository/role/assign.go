package role

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity"
	"github.com/puny-activity/authentication/internal/interr"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) Assign(ctx context.Context, accountID string, role entity.Role) error {
	return r.assign(ctx, r.db, accountID, role)
}

func (r Repository) AssignTx(ctx context.Context, tx *sqlx.Tx, accountID string, role entity.Role) error {
	return r.assign(ctx, tx, accountID, role)
}

func (r Repository) assign(ctx context.Context, queryer queryer.Queryer, accountID string, role entity.Role) error {
	query := `
INSERT INTO account_roles(account_id, role)
VALUES ($1, $2)
`

	_, err := queryer.ExecContext(ctx, query,
		accountID,
		role.Name())
	if err != nil {
		return werr.WrapEE(interr.DatabaseFailedToExecuteQuery, err)
	}

	return nil
}
