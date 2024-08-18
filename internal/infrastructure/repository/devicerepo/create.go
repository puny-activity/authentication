package devicerepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/device"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

type createRepo struct {
	ID          uuid.UUID `db:"id"`
	AccountID   uuid.UUID `db:"account_id"`
	Name        string    `db:"name"`
	Fingerprint string    `db:"fingerprint"`
}

func (r Repository) Create(ctx context.Context, accountID account.ID, deviceToCreate device.Device) error {
	return r.create(ctx, r.db, accountID, deviceToCreate)
}

func (r Repository) CreateTx(ctx context.Context, tx *sqlx.Tx, accountID account.ID, deviceToCreate device.Device) error {
	return r.create(ctx, tx, accountID, deviceToCreate)
}

func (r Repository) create(ctx context.Context, queryer queryer.Queryer, accountID account.ID, deviceToCreate device.Device) error {
	if deviceToCreate.ID == nil {
		return errs.DatabaseUndefinedID
	}

	query := `
INSERT INTO devices(id, account_id, name, fingerprint)
VALUES (:id, :account_id, :name, :fingerprint)
`

	parameter := createRepo{
		ID:          uuid.UUID(*deviceToCreate.ID),
		AccountID:   uuid.UUID(accountID),
		Name:        deviceToCreate.Name,
		Fingerprint: deviceToCreate.Fingerprint,
	}

	_, err := queryer.NamedExecContext(ctx, query, parameter)
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return nil
}
