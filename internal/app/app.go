package app

import (
	"github.com/puny-activity/authentication/internal/config"
	accountrepo "github.com/puny-activity/authentication/internal/infrastructure/database/postgres/repository/account"
	rolerepo "github.com/puny-activity/authentication/internal/infrastructure/database/postgres/repository/role"
	accountuc "github.com/puny-activity/authentication/internal/usecase/account"
	"github.com/puny-activity/authentication/pkg/pstgrs"
	"github.com/puny-activity/authentication/pkg/txmanager"
	"github.com/rs/zerolog"
)

type App struct {
	AccountUseCase *accountuc.UseCase
	log            *zerolog.Logger
}

func New(cfg config.Config, log *zerolog.Logger) *App {
	db, err := pstgrs.New(cfg)
	if err != nil {
		panic(err)
	}

	txManager := txmanager.New(db.DB)

	accountRepo := accountrepo.New(db.DB, txManager, log)
	roleRepo := rolerepo.New(db.DB, txManager, log)

	accountUseCase := accountuc.New(accountRepo, roleRepo, txManager, log)

	return &App{
		AccountUseCase: accountUseCase,
		log:            log,
	}
}
