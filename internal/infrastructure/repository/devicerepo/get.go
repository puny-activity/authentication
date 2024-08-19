package devicerepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/device"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/util"
	"github.com/puny-activity/authentication/pkg/werr"
)

type getRepo struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Fingerprint string    `db:"fingerprint"`
}

func (r Repository) Get(ctx context.Context, deviceID device.ID) (device.Device, error) {
	return r.get(ctx, r.db, deviceID)
}

func (r Repository) GetTx(ctx context.Context, tx *sqlx.Tx, deviceID device.ID) (device.Device, error) {
	return r.get(ctx, tx, deviceID)
}

func (r Repository) get(ctx context.Context, queryer queryer.Queryer, deviceID device.ID) (device.Device, error) {
	query := `
SELECT id,
       name,
       fingerprint
FROM devices
WHERE id = $1
`

	var deviceRepo getRepo
	err := queryer.GetContext(ctx, &deviceRepo, query, uuid.UUID(deviceID))
	if err != nil {
		return device.Device{}, werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	deviceResponse := device.Device{
		ID:          util.ToPointer(device.ID(deviceRepo.ID)),
		Name:        deviceRepo.Name,
		Fingerprint: deviceRepo.Fingerprint,
	}

	return deviceResponse, nil
}
