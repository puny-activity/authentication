package role

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/role"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) Assign(ctx context.Context, accountID uuid.UUID, role role.Role) error {
	return r.assign(ctx, r.db, accountID, role)
}

func (r Repository) AssignTx(ctx context.Context, tx *sqlx.Tx, accountID uuid.UUID, role role.Role) error {
	return r.assign(ctx, tx, accountID, role)
}

func (r Repository) assign(ctx context.Context, queryer queryer.Queryer, accountID uuid.UUID, role role.Role) error {
	query := `
INSERT INTO account_roles(account_id, role)
VALUES ($1, $2)
`

	_, err := queryer.ExecContext(ctx, query,
		accountID.String(),
		role.Name())
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return nil
}
