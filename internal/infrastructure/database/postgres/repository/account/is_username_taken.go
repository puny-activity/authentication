package account

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) IsUsernameTaken(ctx context.Context, username string) (bool, error) {
	return r.isUsernameTaken(ctx, r.db, username)
}

func (r Repository) IsUsernameTakenTx(ctx context.Context, tx *sqlx.Tx, username string) (bool, error) {
	return r.isUsernameTaken(ctx, tx, username)
}

func (r Repository) isUsernameTaken(ctx context.Context, queryer queryer.Queryer, username string) (bool, error) {
	query := `
SELECT EXISTS (SELECT 1
               FROM accounts
               WHERE username = $1)
`

	var isTaken bool

	err := queryer.GetContext(ctx, &isTaken, query,
		username)
	if err != nil {
		return false, werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return isTaken, nil
}
