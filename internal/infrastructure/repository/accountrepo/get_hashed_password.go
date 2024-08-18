package accountrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/account/credential/password"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) GetHashedPassword(ctx context.Context, accountID account.ID) (password.Hashed, error) {
	return r.getHashedPassword(ctx, r.db, accountID)
}

func (r Repository) GetHashedPasswordTx(ctx context.Context, tx *sqlx.Tx, accountID account.ID) (password.Hashed, error) {
	return r.getHashedPassword(ctx, tx, accountID)
}

func (r Repository) getHashedPassword(ctx context.Context, queryer queryer.Queryer, accountID account.ID) (password.Hashed, error) {
	query := `
SELECT a.password_hash
FROM accounts a
WHERE a.id = $1
`

	var hashedPasswordStr string
	err := queryer.GetContext(ctx, &hashedPasswordStr, query, uuid.UUID(accountID))
	if err != nil {
		return password.Hashed{}, werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	hashedPassword := password.NewHashed(hashedPasswordStr)
	return hashedPassword, nil
}
