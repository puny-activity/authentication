package main

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	controllerhealth "github.com/puny-activity/authentication/api/http/healthcheck/controller"
	routerhealth "github.com/puny-activity/authentication/api/http/healthcheck/router"
	"github.com/puny-activity/authentication/api/http/middleware"
	controllerv1 "github.com/puny-activity/authentication/api/http/v1/controller"
	routerv1 "github.com/puny-activity/authentication/api/http/v1/router"
	"github.com/puny-activity/authentication/config"
	"github.com/puny-activity/authentication/internal/app"
	"github.com/puny-activity/authentication/internal/interr"
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

	app := app.New(cfg, log)

	writer := httpresp.NewWriter()

	controllerHealthCheck := controllerhealth.New(writer, log)
	controllerV1 := controllerv1.New(app, writer, log)

	chiMux := chimux.New()

	middleware := middleware.New(log)

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", json.Unmarshal)
	bundle.LoadMessageFile("internal/locale/en-US.json")
	bundle.LoadMessageFile("internal/locale/ru-RU.json")

	localizer := loclzr.New(bundle)

	errorStorage := interr.NewErrorStorage(localizer)

	wrapper := httpresp.NewWrapper(writer, errorStorage, log)

	routerHealthCheck := routerhealth.New(cfg, chiMux, middleware, wrapper, controllerHealthCheck, log)
	routerV1 := routerv1.New(cfg, chiMux, middleware, wrapper, controllerV1, log)

	routerHealthCheck.Setup()
	routerV1.Setup()

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
