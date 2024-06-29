package accesstoken

import "github.com/google/uuid"

type Content struct {
	AccountID      uuid.UUID
	DeviceID       uuid.UUID
	RefreshTokenID uuid.UUID
	Roles          []string
	IssuedAt       int64
	ExpiresAt      int64
}

type Payload struct {
	AccountID      string   `json:"accountId"`
	DeviceID       string   `json:"deviceId"`
	RefreshTokenID string   `json:"refreshTokenId"`
	Roles          []string `json:"roles"`
	IssuedAt       int64    `json:"issuedAt"`
	ExpiresAt      int64    `json:"expiresAt"`
}

func (e *Content) ToPayload() Payload {
	return Payload{
		AccountID:      e.AccountID.String(),
		DeviceID:       e.DeviceID.String(),
		RefreshTokenID: e.RefreshTokenID.String(),
		Roles:          e.Roles,
		IssuedAt:       e.IssuedAt,
		ExpiresAt:      e.ExpiresAt,
	}
}
