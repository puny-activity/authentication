package accesstoken

import (
	"errors"
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/account/credential/email"
	"github.com/puny-activity/authentication/internal/entity/device"
	"github.com/puny-activity/authentication/internal/entity/role"
	"github.com/puny-activity/authentication/internal/entity/token/refreshtoken"
	"github.com/puny-activity/authentication/pkg/util"
)

type ID uuid.UUID

type AccessToken struct {
	Base
	Payload
}

type Base struct {
	ID             *ID
	IssuedAt       carbon.Carbon
	ExpiresAt      carbon.Carbon
	RefreshTokenID refreshtoken.ID
}

type Payload struct {
	AccountID   account.ID
	Email       email.Email
	Nickname    string
	Role        role.Role
	DeviceID    device.ID
	DeviceName  string
	Fingerprint string
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
		AccountID:   *account.ID,
		Email:       account.Email,
		Nickname:    account.Nickname,
		Role:        account.Role,
		DeviceID:    *device.ID,
		DeviceName:  device.Name,
		Fingerprint: device.Fingerprint,
	}, nil
}
