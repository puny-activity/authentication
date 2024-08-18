package accesstokenservice

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
	"github.com/puny-activity/authentication/internal/entity/token/accesstoken"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (s Service) Encode(token accesstoken.AccessToken) (string, error) {
	claims := jwt.MapClaims{
		"jti":               uuid.UUID(*token.Base.ID),
		"iat":               token.Base.IssuedAt.ToStdTime().Unix(),
		"exp":               token.Base.ExpiresAt.ToStdTime().Unix(),
		"sub":               uuid.UUID(token.Payload.AccountID),
		"deviceId":          uuid.UUID(token.Payload.DeviceID),
		"role":              token.Payload.Role.Name(),
		"nickname":          token.Payload.Nickname,
		"deviceName":        token.Payload.DeviceName,
		"deviceFingerprint": token.Payload.Fingerprint,
	}

	refreshTokenByte := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := refreshTokenByte.SignedString([]byte(s.cfg.SecretKey()))
	if err != nil {
		return "", werr.WrapSE("failed to sign refresh token", err)
	}

	return refreshToken, nil
}
