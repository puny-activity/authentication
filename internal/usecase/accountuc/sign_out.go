package accountuc

import (
	"context"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (u *UseCase) SignOut(ctx context.Context, refreshTokenString string) error {
	oldRefreshToken, err := u.refreshTokenService.Decode(refreshTokenString)
	if err != nil {
		return werr.WrapSE("failed to decode refresh token", err)
	}

	err = u.refreshTokenRepo.Delete(ctx, *oldRefreshToken.ID)
	if err != nil {
		return werr.WrapSE("failed to delete old refresh token", err)
	}

	return nil
}
