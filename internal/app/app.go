package app

import (
	"github.com/puny-activity/authentication/internal/config"
	accountrepo "github.com/puny-activity/authentication/internal/infrastructure/database/postgres/repository/account"
	accountuc "github.com/puny-activity/authentication/internal/usecase/account"
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

func New(cfg config.Config, log *zerolog.Logger) *App {
	db, err := database.New(cfg)
	if err != nil {
		panic(err)
	}
	err = db.RunMigrations()
	if err != nil {
		panic(err)
	}

	txManager := txmanager.New(db.DB)

	accountRepo := accountrepo.New(db.DB, txManager, log)

	accountUseCase := accountuc.New(accountRepo, txManager, log)

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
