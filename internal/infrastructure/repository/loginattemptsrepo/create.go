package loginattemptsrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/loginattempt"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

type createRepo struct {
	ID          uuid.UUID `db:"id"`
	Email       string    `db:"email"`
	Success     bool      `db:"success"`
	AttemptedAt string    `db:"attempted_at"`
}

func (r Repository) Create(ctx context.Context, row loginattempt.Row) error {
	return r.create(ctx, r.db, row)
}

func (r Repository) CreateTx(ctx context.Context, tx *sqlx.Tx, row loginattempt.Row) error {
	return r.create(ctx, tx, row)
}

func (r Repository) create(ctx context.Context, queryer queryer.Queryer, row loginattempt.Row) error {
	if row.ID == nil {
		return errs.DatabaseUndefinedID
	}

	query := `
INSERT INTO login_attempts(id, email, success, attempted_at)
VALUES (:id, :email, :success, :attempted_at)
`

	parameter := createRepo{
		ID:          uuid.UUID(*row.ID),
		Email:       row.Email.String(),
		Success:     row.Success,
		AttemptedAt: row.AttemptedAt.ToDateTimeString(),
	}

	_, err := queryer.NamedExecContext(ctx, query, parameter)
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return nil
}
