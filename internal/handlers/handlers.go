package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/jetqin/goapi/internal/middleware"
	"github.com/sirupsen/logrus"
)

func Handlers(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)
	r.Use(middleware.Logger(logrus.StandardLogger()))
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)

	})
}
