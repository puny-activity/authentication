package accountuc

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/token"
	"github.com/puny-activity/authentication/internal/entity/token/accesstoken"
	"github.com/puny-activity/authentication/internal/entity/token/refreshtoken"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (u *UseCase) Refresh(ctx context.Context, oldRefreshTokenString string) (token.Pair, error) {
	oldRefreshToken, err := u.refreshTokenService.Decode(oldRefreshTokenString)
	if err != nil {
		return token.Pair{}, werr.WrapSE("failed to decode refresh token", err)
	}

	var refreshTokenString string
	var accessTokenString string

	err = u.txManager.Transaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		err := u.refreshTokenRepo.DeleteTx(ctx, tx, *oldRefreshToken.ID)
		if err != nil {
			return werr.WrapSE("failed to delete old refresh token", err)
		}

		targetAccount, err := u.accountRepo.GetTx(ctx, tx, oldRefreshToken.AccountID)
		if err != nil {
			return werr.WrapSE("failed to get account", err)
		}
		targetDevice, err := u.deviceRepo.GetTx(ctx, tx, oldRefreshToken.DeviceID)
		if err != nil {
			return werr.WrapSE("failed to get device", err)
		}

		refreshTokenPayload, err := refreshtoken.NewPayload(targetAccount, targetDevice)
		if err != nil {
			return werr.WrapSE("failed to create refresh token payload", err)
		}
		refreshToken, err := u.refreshTokenService.Generate(refreshTokenPayload)
		if err != nil {
			return werr.WrapSE("failed to generate refresh token payload", err)
		}
		refreshTokenString, err = u.refreshTokenService.Encode(refreshToken)
		err = u.refreshTokenRepo.CreateTx(ctx, tx, *targetDevice.ID, refreshToken.Base)
		if err != nil {
			return werr.WrapSE("failed to create refresh token", err)
		}

		accessTokenPayload, err := accesstoken.NewPayload(targetAccount, targetDevice)
		if err != nil {
			return werr.WrapSE("failed to create access token payload", err)
		}
		accessToken, err := u.accessTokenService.Generate(*refreshToken.ID, accessTokenPayload)
		if err != nil {
			return werr.WrapSE("failed to generate access token", err)
		}
		accessTokenString, err = u.accessTokenService.Encode(accessToken)

		return nil
	})
	if err != nil {
		return token.Pair{}, err
	}

	return token.Pair{
		RefreshToken: refreshTokenString,
		AccessToken:  accessTokenString,
	}, nil
}
