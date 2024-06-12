package controller

import (
	"github.com/puny-activity/authentication/internal/app"
	"github.com/puny-activity/authentication/pkg/httpresp"
	"github.com/rs/zerolog"
)

type Controller struct {
	app            *app.App
	responseWriter *httpresp.Writer
	log            *zerolog.Logger
}

func New(app *app.App, responseWriter *httpresp.Writer, log *zerolog.Logger) *Controller {
	return &Controller{
		app:            app,
		responseWriter: responseWriter,
		log:            log,
	}
}
