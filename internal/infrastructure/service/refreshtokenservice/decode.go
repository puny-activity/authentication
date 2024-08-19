package refreshtokenservice

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/device"
	"github.com/puny-activity/authentication/internal/entity/token/refreshtoken"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/util"
	"github.com/puny-activity/authentication/pkg/werr"
	"time"
)

func (s Service) Decode(tokenString string) (refreshtoken.RefreshToken, error) {
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.SecretKey()), nil
	})
	if err != nil {
		return refreshtoken.RefreshToken{}, werr.WrapES(errs.InvalidRefreshToken, err.Error())
	}
	if !parsedToken.Valid {
		return refreshtoken.RefreshToken{}, errs.InvalidRefreshToken
	}

	issuedAtUnix := int64(claims["iat"].(float64))
	expiresAtUnix := int64(claims["exp"].(float64))
	issuedAt := carbon.Time2Carbon(time.Unix(issuedAtUnix, 0))
	expiresAt := carbon.Time2Carbon(time.Unix(expiresAtUnix, 0))

	if issuedAt.Gt(carbon.Now()) || expiresAt.Lt(carbon.Now()) {
		return refreshtoken.RefreshToken{}, errs.InvalidRefreshToken
	}

	refreshToken := refreshtoken.RefreshToken{}
	refreshToken.Base.IssuedAt = issuedAt
	refreshToken.Base.ExpiresAt = expiresAt
	jtiUUID, err := uuid.Parse(claims["jti"].(string))
	if err != nil {
		return refreshtoken.RefreshToken{}, werr.WrapES(errs.InvalidRefreshToken, err.Error())
	}
	refreshToken.Base.ID = util.ToPointer(refreshtoken.ID(jtiUUID))
	subUUID, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return refreshtoken.RefreshToken{}, werr.WrapES(errs.InvalidRefreshToken, err.Error())
	}
	refreshToken.Payload.AccountID = account.ID(subUUID)
	deviceUUID, err := uuid.Parse(claims["deviceId"].(string))
	if err != nil {
		return refreshtoken.RefreshToken{}, werr.WrapES(errs.InvalidRefreshToken, err.Error())
	}
	refreshToken.Payload.DeviceID = device.ID(deviceUUID)

	return refreshToken, nil
}
