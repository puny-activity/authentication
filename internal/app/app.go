package app

import (
	"github.com/puny-activity/authentication/config"
	"github.com/puny-activity/authentication/internal/infrastructure/repository/accountrepo"
	"github.com/puny-activity/authentication/internal/infrastructure/repository/devicerepo"
	"github.com/puny-activity/authentication/internal/infrastructure/repository/loginattemptsrepo"
	"github.com/puny-activity/authentication/internal/infrastructure/repository/refreshtokenrepo"
	"github.com/puny-activity/authentication/internal/infrastructure/service/accesstokenservice"
	"github.com/puny-activity/authentication/internal/infrastructure/service/refreshtokenservice"
	accountuc "github.com/puny-activity/authentication/internal/usecase/accountuc"
	"github.com/puny-activity/authentication/pkg/database"
	"github.com/puny-activity/authentication/pkg/txmanager"
	"github.com/puny-activity/authentication/pkg/werr"
	"github.com/rs/zerolog"
)

type App struct {
	AccountUseCase *accountuc.UseCase
	db             *database.Database
	log            *zerolog.Logger
}

func New(cfg *config.Config, log *zerolog.Logger) *App {
	db, err := database.New(cfg.DB)
	if err != nil {
		panic(err)
	}
	err = db.RunMigrations()
	if err != nil {
		panic(err)
	}

	txManager := txmanager.New(db.DB)

	accountRepo := accountrepo.New(db.DB, txManager, log)
	deviceRepo := devicerepo.New(db.DB, txManager, log)
	refreshTokenRepo := refreshtokenrepo.New(db.DB, txManager, log)
	loginAttemptsRepo := loginattemptsrepo.New(db.DB, txManager, log)

	refreshTokenService := refreshtokenservice.New(cfg.App.RefreshToken)
	accessTokenService := accesstokenservice.New(cfg.App.AccessToken)

	accountUseCase := accountuc.New(accountRepo, deviceRepo, refreshTokenRepo, loginAttemptsRepo, refreshTokenService,
		accessTokenService, txManager, log)

	return &App{
		AccountUseCase: accountUseCase,
		db:             db,
		log:            log,
	}
}

func (a *App) Close() error {
	err := a.db.Close()
	if err != nil {
		return werr.WrapSE("failed to close database", err)
	}

	return nil
}
