package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/puny-activity/authentication/api/http/main/controller"
	"github.com/puny-activity/authentication/api/http/middleware"
	"github.com/puny-activity/authentication/config"
	"github.com/puny-activity/authentication/pkg/httpresp"
	"github.com/rs/zerolog"
)

type Router struct {
	cfg        *config.Config
	router     *chi.Mux
	middleware *middleware.Middleware
	wrapper    *httpresp.Wrapper
	controller *controller.Controller
	log        *zerolog.Logger
}

func New(cfg *config.Config, router *chi.Mux, middleware *middleware.Middleware, wrapper *httpresp.Wrapper,
	controller *controller.Controller, log *zerolog.Logger) *Router {
	return &Router{
		cfg:        cfg,
		router:     router,
		middleware: middleware,
		wrapper:    wrapper,
		controller: controller,
		log:        log,
	}
}

func (r *Router) Setup() {
	r.router.Group(func(router chi.Router) {
		router.Use(r.middleware.AcceptLanguageMiddleware)

		router.Route("/sign-up", func(router chi.Router) {
			router.Post("/", r.wrapper.WrapAnonymous(r.controller.SignUp))
		})

		router.Route("/sign-in", func(router chi.Router) {
			router.Post("/", r.wrapper.WrapAnonymous(r.controller.SignIn))
		})

		router.Route("/refresh", func(router chi.Router) {
			router.Post("/", r.wrapper.WrapAnonymous(r.controller.Refresh))
		})

		router.Route("/sign-out", func(router chi.Router) {
			router.Post("/", r.wrapper.WrapAnonymous(r.controller.SignOut))
		})
	})
}
