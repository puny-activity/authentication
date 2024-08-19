package refreshtokenrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/token/refreshtoken"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) Delete(ctx context.Context, refreshTokenID refreshtoken.ID) error {
	return r.delete(ctx, r.db, refreshTokenID)
}

func (r Repository) DeleteTx(ctx context.Context, tx *sqlx.Tx, refreshTokenID refreshtoken.ID) error {
	return r.delete(ctx, tx, refreshTokenID)
}

func (r Repository) delete(ctx context.Context, queryer queryer.Queryer, refreshTokenID refreshtoken.ID) error {
	query := `
DELETE
FROM refresh_tokens rt
WHERE rt.id = $1
`

	res, err := queryer.ExecContext(ctx, query, uuid.UUID(refreshTokenID))
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}
	if rowsAffected == 0 {
		return errs.RefreshTokenNotFound
	}

	return nil
}
