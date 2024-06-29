package refreshtoken

import "github.com/google/uuid"

type Payload struct {
	RefreshTokenID *uuid.UUID `json:"refreshTokenId"`
	IssuedAt       int64      `json:"issuedAt"`
	ExpiresAt      int64      `json:"expiresAt"`
}
