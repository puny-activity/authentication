package accountuc

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/account/credential/email"
	"github.com/puny-activity/authentication/internal/entity/account/credential/password"
	"github.com/puny-activity/authentication/internal/entity/device"
	"github.com/puny-activity/authentication/internal/entity/loginattempt"
	"github.com/puny-activity/authentication/internal/entity/token/accesstoken"
	"github.com/puny-activity/authentication/internal/entity/token/refreshtoken"
	"github.com/puny-activity/authentication/pkg/txmanager"
	"github.com/rs/zerolog"
)

type UseCase struct {
	accountRepo         accountRepository
	deviceRepo          deviceRepository
	refreshTokenRepo    refreshTokenRepository
	loginAttemptsRepo   loginAttemptsRepository
	refreshTokenService refreshTokenService
	accessTokenService  accessTokenService
	txManager           txmanager.Transactor
	log                 *zerolog.Logger
}

func New(accountRepo accountRepository, deviceRepo deviceRepository, refreshTokenRepo refreshTokenRepository,
	loginAttemptsRepo loginAttemptsRepository, refreshTokenService refreshTokenService,
	accessTokenService accessTokenService, txManager txmanager.Transactor, log *zerolog.Logger) *UseCase {
	return &UseCase{
		accountRepo:         accountRepo,
		deviceRepo:          deviceRepo,
		refreshTokenRepo:    refreshTokenRepo,
		loginAttemptsRepo:   loginAttemptsRepo,
		refreshTokenService: refreshTokenService,
		accessTokenService:  accessTokenService,
		txManager:           txManager,
		log:                 log,
	}
}

type accountRepository interface {
	IsEmailTakenTx(ctx context.Context, tx *sqlx.Tx, email email.Email) (bool, error)
	IsNicknameTakenTx(ctx context.Context, tx *sqlx.Tx, nickname string) (bool, error)
	CreateTx(ctx context.Context, tx *sqlx.Tx, accountToCreate account.Account, hashedPasswordToCreate password.Hashed) error
	CountTx(ctx context.Context, tx *sqlx.Tx) (int, error)
	GetByEmailTx(ctx context.Context, tx *sqlx.Tx, targetEmail email.Email) (account.Account, error)
	GetHashedPasswordTx(ctx context.Context, tx *sqlx.Tx, accountID account.ID) (password.Hashed, error)
	GetTx(ctx context.Context, tx *sqlx.Tx, accountID account.ID) (account.Account, error)
}

type deviceRepository interface {
	DeleteIfExistsByFingerprintTx(ctx context.Context, tx *sqlx.Tx, fingerprint string) error
	CreateTx(ctx context.Context, tx *sqlx.Tx, accountID account.ID, deviceToCreate device.Device) error
	GetTx(ctx context.Context, tx *sqlx.Tx, deviceID device.ID) (device.Device, error)
}

type refreshTokenRepository interface {
	DeleteTx(ctx context.Context, tx *sqlx.Tx, refreshTokenID refreshtoken.ID) error
	DeleteIfExistsByDeviceFingerprintTx(ctx context.Context, tx *sqlx.Tx, fingerprint string) error
	CreateTx(ctx context.Context, tx *sqlx.Tx, deviceID device.ID, baseToken refreshtoken.Base) error
	Delete(ctx context.Context, refreshTokenID refreshtoken.ID) error
}

type loginAttemptsRepository interface {
	Create(ctx context.Context, row loginattempt.Row) error
}

type refreshTokenService interface {
	Generate(payload refreshtoken.Payload) (refreshtoken.RefreshToken, error)
	Encode(token refreshtoken.RefreshToken) (string, error)
	Decode(tokenString string) (refreshtoken.RefreshToken, error)
}

type accessTokenService interface {
	Generate(parentRefreshTokenID refreshtoken.ID, payload accesstoken.Payload) (accesstoken.AccessToken, error)
	Encode(token accesstoken.AccessToken) (string, error)
}
