package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/kaizakin/goapi/internal/middleware"
)

func Handler(r *chi.Mux){
	r.Use(chimiddle.StripSlashes) // global middleware to strip the trailing slashes in the url which might cause issue.

	r.Route("/account", func (router chi.Router){
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}