package refreshtokenrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/device"
	"github.com/puny-activity/authentication/internal/entity/token/refreshtoken"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

type createRepo struct {
	ID        uuid.UUID `db:"id"`
	DeviceID  uuid.UUID `db:"device_id"`
	IssuedAt  string    `db:"issued_at"`
	ExpiresAt string    `db:"expires_at"`
}

func (r Repository) Create(ctx context.Context, deviceID device.ID, baseToken refreshtoken.Base) error {
	return r.create(ctx, r.db, deviceID, baseToken)
}

func (r Repository) CreateTx(ctx context.Context, tx *sqlx.Tx, deviceID device.ID, baseToken refreshtoken.Base) error {
	return r.create(ctx, tx, deviceID, baseToken)
}

func (r Repository) create(ctx context.Context, queryer queryer.Queryer, deviceID device.ID, baseToken refreshtoken.Base) error {
	if baseToken.ID == nil {
		return errs.DatabaseUndefinedID
	}

	query := `
INSERT INTO refresh_tokens(id, device_id, issued_at, expires_at)
VALUES (:id, :device_id, :issued_at, :expires_at)
`

	parameter := createRepo{
		ID:        uuid.UUID(*baseToken.ID),
		DeviceID:  uuid.UUID(deviceID),
		IssuedAt:  baseToken.IssuedAt.ToDateTimeString(),
		ExpiresAt: baseToken.ExpiresAt.ToDateTimeString(),
	}

	_, err := queryer.NamedExecContext(ctx, query, parameter)
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return nil
}
