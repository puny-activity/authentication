package entity

type RefreshTokenPayload struct {
	RefreshTokenID string `json:"refreshTokenId"`
	IssuedAt       int64  `json:"issuedAt"`
	ExpiresAt      int64  `json:"expiresAt"`
}
