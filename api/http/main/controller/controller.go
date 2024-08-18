package controller

import (
	"github.com/puny-activity/authentication/config"
	"github.com/puny-activity/authentication/internal/app"
	"github.com/puny-activity/authentication/pkg/httpresp"
	"github.com/rs/zerolog"
)

type Controller struct {
	cfg            *config.Config
	app            *app.App
	responseWriter *httpresp.Writer
	log            *zerolog.Logger
}

func New(cfg *config.Config, app *app.App, responseWriter *httpresp.Writer, log *zerolog.Logger) *Controller {
	return &Controller{
		cfg:            cfg,
		app:            app,
		responseWriter: responseWriter,
		log:            log,
	}
}
