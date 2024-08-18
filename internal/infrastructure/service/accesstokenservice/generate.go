package accesstokenservice

import (
	"github.com/golang-module/carbon"
	"github.com/puny-activity/authentication/internal/entity/token/accesstoken"
)

func (s Service) Generate(payload accesstoken.Payload) (accesstoken.AccessToken, error) {
	issuedAt := carbon.Now()
	expiresAt := carbon.Now().AddSeconds(s.cfg.TTLSecond())

	baseToken := accesstoken.Base{
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}
	baseToken = baseToken.GenerateID()

	token := accesstoken.AccessToken{
		Base:    baseToken,
		Payload: payload,
	}

	return token, nil
}
