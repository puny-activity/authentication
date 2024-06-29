package account

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) AccountsCount(ctx context.Context) (int, error) {
	return r.accountsCount(ctx, r.db)
}

func (r Repository) AccountsCountTx(ctx context.Context, tx *sqlx.Tx) (int, error) {
	return r.accountsCount(ctx, tx)
}

func (r Repository) accountsCount(ctx context.Context, queryer queryer.Queryer) (int, error) {
	query := `
SELECT COUNT(*)
FROM accounts
`

	var count int

	err := queryer.GetContext(ctx, &count, query)
	if err != nil {
		return 0, werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return count, nil
}
