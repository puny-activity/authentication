package devicerepo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) DeleteIfExistsByFingerprint(ctx context.Context, fingerprint string) error {
	return r.deleteIfExistsByFingerprint(ctx, r.db, fingerprint)
}

func (r Repository) DeleteIfExistsByFingerprintTx(ctx context.Context, tx *sqlx.Tx, fingerprint string) error {
	return r.deleteIfExistsByFingerprint(ctx, tx, fingerprint)
}

func (r Repository) deleteIfExistsByFingerprint(ctx context.Context, queryer queryer.Queryer, fingerprint string) error {
	query := `
DELETE FROM devices d
WHERE d.fingerprint = $1
`

	_, err := queryer.ExecContext(ctx, query, fingerprint)
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return nil
}
