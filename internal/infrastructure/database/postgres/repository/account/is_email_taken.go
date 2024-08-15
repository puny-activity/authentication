package account

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	return r.isEmailTaken(ctx, r.db, email)
}

func (r Repository) IsEmailTakenTx(ctx context.Context, tx *sqlx.Tx, email string) (bool, error) {
	return r.isEmailTaken(ctx, tx, email)
}

func (r Repository) isEmailTaken(ctx context.Context, queryer queryer.Queryer, email string) (bool, error) {
	query := `
SELECT EXISTS (SELECT 1
               FROM accounts
               WHERE email = $1)
`

	var isTaken bool

	err := queryer.GetContext(ctx, &isTaken, query,
		email)
	if err != nil {
		return false, werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return isTaken, nil
}
