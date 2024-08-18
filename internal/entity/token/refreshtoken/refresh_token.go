package refreshtoken

import (
	"errors"
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/device"
	"github.com/puny-activity/authentication/pkg/util"
)

type ID uuid.UUID

type RefreshToken struct {
	Base
	Payload
}

type Base struct {
	ID        *ID
	IssuedAt  carbon.Carbon
	ExpiresAt carbon.Carbon
}

type Payload struct {
	AccountID account.ID
	DeviceID  device.ID
}

func (e Base) GenerateID() Base {
	e.ID = util.ToPointer(ID(uuid.New()))
	return e
}

func NewPayload(account account.Account, device device.Device) (Payload, error) {
	if account.ID == nil || device.ID == nil {
		return Payload{}, errors.New("invalid payload data")
	}

	return Payload{
		AccountID: *account.ID,
		DeviceID:  *device.ID,
	}, nil
}
