package accesstokenservice

import (
	"github.com/golang-module/carbon"
	"github.com/puny-activity/authentication/internal/entity/token/accesstoken"
	"github.com/puny-activity/authentication/internal/entity/token/refreshtoken"
)

func (s Service) Generate(parentRefreshTokenID refreshtoken.ID, payload accesstoken.Payload) (accesstoken.AccessToken, error) {
	issuedAt := carbon.Now()
	expiresAt := carbon.Now().AddSeconds(s.cfg.TTLSecond())

	baseToken := accesstoken.Base{
		IssuedAt:       issuedAt,
		ExpiresAt:      expiresAt,
		RefreshTokenID: parentRefreshTokenID,
	}
	baseToken = baseToken.GenerateID()

	token := accesstoken.AccessToken{
		Base:    baseToken,
		Payload: payload,
	}

	return token, nil
}
