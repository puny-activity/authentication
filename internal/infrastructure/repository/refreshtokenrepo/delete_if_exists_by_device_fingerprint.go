package refreshtokenrepo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (r Repository) DeleteIfExistsByDeviceFingerprint(ctx context.Context, fingerprint string) error {
	return r.deleteIfExistsByDeviceFingerprint(ctx, r.db, fingerprint)
}

func (r Repository) DeleteIfExistsByDeviceFingerprintTx(ctx context.Context, tx *sqlx.Tx, fingerprint string) error {
	return r.deleteIfExistsByDeviceFingerprint(ctx, tx, fingerprint)
}

func (r Repository) deleteIfExistsByDeviceFingerprint(ctx context.Context, queryer queryer.Queryer, fingerprint string) error {
	query := `
DELETE
FROM refresh_tokens rt
WHERE rt.id IN (SELECT rt.id
                FROM devices d
                         JOIN refresh_tokens rt ON d.id = rt.device_id
                WHERE d.fingerprint = $1)
`

	_, err := queryer.ExecContext(ctx, query, fingerprint)
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return nil
}
