package accountrepo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account/credential/email"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) IsEmailTaken(ctx context.Context, targetEmail email.Email) (bool, error) {
	return r.isEmailTaken(ctx, r.db, targetEmail)
}

func (r Repository) IsEmailTakenTx(ctx context.Context, tx *sqlx.Tx, targetEmail email.Email) (bool, error) {
	return r.isEmailTaken(ctx, tx, targetEmail)
}

func (r Repository) isEmailTaken(ctx context.Context, queryer queryer.Queryer, targetEmail email.Email) (bool, error) {
	query := `
SELECT EXISTS (SELECT 1
               FROM accounts
               WHERE email = $1)
`

	var isTaken bool

	err := queryer.GetContext(ctx, &isTaken, query,
		targetEmail.String())
	if err != nil {
		return false, werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return isTaken, nil
}
