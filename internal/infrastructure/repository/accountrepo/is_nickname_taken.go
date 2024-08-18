package accountrepo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) IsNicknameTaken(ctx context.Context, nickname string) (bool, error) {
	return r.isNicknameTaken(ctx, r.db, nickname)
}

func (r Repository) IsNicknameTakenTx(ctx context.Context, tx *sqlx.Tx, nickname string) (bool, error) {
	return r.isNicknameTaken(ctx, tx, nickname)
}

func (r Repository) isNicknameTaken(ctx context.Context, queryer queryer.Queryer, nickname string) (bool, error) {
	query := `
SELECT EXISTS (SELECT 1
               FROM accounts
               WHERE nickname = $1)
`

	var isTaken bool

	err := queryer.GetContext(ctx, &isTaken, query,
		nickname)
	if err != nil {
		return false, werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return isTaken, nil
}
