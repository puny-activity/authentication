package entity

type AccessTokenPayload struct {
	AccountID      string   `json:"accountId"`
	DeviceID       string   `json:"deviceId"`
	Roles          []string `json:"roles"`
	RefreshTokenID string   `json:"refreshTokenId"`
	IssuedAt       int64    `json:"issuedAt"`
	ExpiresAt      int64    `json:"expiresAt"`
}
