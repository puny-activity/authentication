package app

import (
	"github.com/puny-activity/authentication/internal/config"
	"github.com/puny-activity/authentication/pkg/pstgrs"
	"github.com/rs/zerolog"
)

type App struct {
	log *zerolog.Logger
}

func New(cfg config.Config, log *zerolog.Logger) *App {
	db, err := pstgrs.New(cfg)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to close database")
		}
	}()

	return &App{
		log: log,
	}
}
