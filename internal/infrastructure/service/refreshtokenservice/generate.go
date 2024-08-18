package refreshtokenservice

import (
	"github.com/golang-module/carbon"
	"github.com/puny-activity/authentication/internal/entity/token/refreshtoken"
)

func (s Service) Generate(payload refreshtoken.Payload) (refreshtoken.RefreshToken, error) {
	issuedAt := carbon.Now()
	expiresAt := carbon.Now().AddSeconds(s.cfg.TTLSecond())

	baseToken := refreshtoken.Base{
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}
	baseToken = baseToken.GenerateID()

	token := refreshtoken.RefreshToken{
		Base:    baseToken,
		Payload: payload,
	}

	return token, nil
}
