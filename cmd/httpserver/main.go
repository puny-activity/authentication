package main

import (
	"encoding/json"
	"github.com/golang-module/carbon"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	controllerhealth "github.com/puny-activity/authentication/api/http/healthcheck/controller"
	routerhealth "github.com/puny-activity/authentication/api/http/healthcheck/router"
	"github.com/puny-activity/authentication/api/http/main/controller"
	"github.com/puny-activity/authentication/api/http/main/router"
	"github.com/puny-activity/authentication/api/http/middleware"
	"github.com/puny-activity/authentication/config"
	"github.com/puny-activity/authentication/internal/app"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/internal/locale"
	"github.com/puny-activity/authentication/pkg/chimux"
	"github.com/puny-activity/authentication/pkg/httpresp"
	"github.com/puny-activity/authentication/pkg/httpsrvr"
	"github.com/puny-activity/authentication/pkg/loclzr"
	"github.com/puny-activity/authentication/pkg/zerologger"
	"golang.org/x/text/language"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		panic(err)
	}

	log, err := zerologger.NewLogger(cfg.Logger.Level)
	if err != nil {
		panic(err)
	}

	carbon.SetTimezone("UTC")

	application := app.New(cfg, log)
	defer func(application *app.App) {
		err := application.Close()
		if err != nil {
			log.Warn().Err(err).Msg("Failed to close application")
		}
	}(application)

	writer := httpresp.NewWriter()

	controllerHealthCheck := controllerhealth.New(writer, log)
	controllerMain := controller.New(cfg, application, writer, log)

	chiMux := chimux.New()

	middlewares := middleware.New(log)

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", json.Unmarshal)
	for _, lang := range locale.Languages {
		_, err := bundle.LoadMessageFile("internal/locale/" + lang.String() + ".json")
		if err != nil {
			log.Warn().Err(err).Msg("Failed to load locale file")
		}
	}

	localizer := loclzr.New(bundle, locale.Languages)

	errorStorage := errs.NewStorage(localizer)

	wrapper := httpresp.NewWrapper(writer, errorStorage, log)

	routerHealthCheck := routerhealth.New(cfg, chiMux, middlewares, wrapper, controllerHealthCheck, log)
	routerMain := router.New(cfg, chiMux, middlewares, wrapper, controllerMain, log)

	routerHealthCheck.Setup()
	routerMain.Setup()

	httpServer := httpsrvr.New(
		chiMux,
		httpsrvr.Addr(cfg.HTTP.Host, cfg.HTTP.Port),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Str("signal", s.String()).Msg("interrupt")
	}

	err = httpServer.Shutdown()
	if err != nil {
		panic(err)
	}
}
