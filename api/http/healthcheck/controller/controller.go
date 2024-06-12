package controller

import (
	"github.com/puny-activity/authentication/pkg/httpresp"
	"github.com/rs/zerolog"
)

type Controller struct {
	responseWriter *httpresp.Writer
	log            *zerolog.Logger
}

func New(responseWriter *httpresp.Writer, log *zerolog.Logger) *Controller {
	return &Controller{
		responseWriter: responseWriter,
		log:            log,
	}
}
