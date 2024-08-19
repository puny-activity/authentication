package accountuc

import (
	"context"
	"github.com/golang-module/carbon"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account/credential"
	"github.com/puny-activity/authentication/internal/entity/account/credential/email"
	"github.com/puny-activity/authentication/internal/entity/account/credential/password"
	"github.com/puny-activity/authentication/internal/entity/device"
	"github.com/puny-activity/authentication/internal/entity/loginattempt"
	"github.com/puny-activity/authentication/internal/entity/token"
	"github.com/puny-activity/authentication/internal/entity/token/accesstoken"
	"github.com/puny-activity/authentication/internal/entity/token/refreshtoken"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (u *UseCase) SignIn(ctx context.Context, credentials credential.Credential, sourceDevice device.Device) (token.Pair, error) {
	var refreshTokenString string
	var accessTokenString string

	err := u.txManager.Transaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		targetAccount, err := u.accountRepo.GetByEmailTx(ctx, tx, credentials.Email)
		if err != nil {
			return werr.WrapSE("failed to get account by email", err)
		}

		accountsHashedPassword, err := u.accountRepo.GetHashedPasswordTx(ctx, tx, *targetAccount.ID)
		if err != nil {
			return werr.WrapSE("failed to get hashed password", err)
		}

		isPasswordMatch := password.IsMatch(credentials.Password, accountsHashedPassword)
		if !isPasswordMatch {
			return errs.WrongPassword
		}

		err = u.refreshTokenRepo.DeleteIfExistsByDeviceFingerprintTx(ctx, tx, sourceDevice.Fingerprint)
		if err != nil {
			return werr.WrapSE("failed to delete refresh token", err)
		}
		err = u.deviceRepo.DeleteIfExistsByFingerprintTx(ctx, tx, sourceDevice.Fingerprint)
		if err != nil {
			return werr.WrapSE("failed to delete device", err)
		}

		deviceToCreate := device.Device{
			Name:        sourceDevice.Name,
			Fingerprint: sourceDevice.Fingerprint,
		}
		deviceToCreate = deviceToCreate.GenerateID()
		err = u.deviceRepo.CreateTx(ctx, tx, *targetAccount.ID, deviceToCreate)
		if err != nil {
			return werr.WrapSE("failed to create device", err)
		}

		refreshTokenPayload, err := refreshtoken.NewPayload(targetAccount, deviceToCreate)
		if err != nil {
			return werr.WrapSE("failed to create refresh token payload", err)
		}
		refreshToken, err := u.refreshTokenService.Generate(refreshTokenPayload)
		if err != nil {
			return werr.WrapSE("failed to generate refresh token payload", err)
		}
		refreshTokenString, err = u.refreshTokenService.Encode(refreshToken)
		err = u.refreshTokenRepo.CreateTx(ctx, tx, *deviceToCreate.ID, refreshToken.Base)
		if err != nil {
			return werr.WrapSE("failed to create refresh token", err)
		}

		accessTokenPayload, err := accesstoken.NewPayload(targetAccount, deviceToCreate)
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
		u.createSignInAttempt(ctx, credentials.Email, false)
		return token.Pair{}, err
	}

	u.createSignInAttempt(ctx, credentials.Email, true)
	return token.Pair{
		RefreshToken: refreshTokenString,
		AccessToken:  accessTokenString,
	}, nil
}

func (u *UseCase) createSignInAttempt(ctx context.Context, email email.Email, success bool) {
	row := loginattempt.Row{
		Email:       email,
		Success:     success,
		AttemptedAt: carbon.Now(),
	}
	row = row.GenerateID()
	err := u.loginAttemptsRepo.Create(ctx, row)
	if err != nil {
		u.log.Error().Err(err).Str("email", email.String()).Bool("success", success).Msg("Failed to create sign in attempt")
	}
}
