package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/hemanG-G/GO_MINI_PROJECTS/tree/main/BasicAPI/internal/middleware"
)

func Handler(r *chi.Mux) { // "H" capital , for public function

	//Global middleware
	r.Use(chimiddle.StripSlashes) // strip slashes at end of URL

	r.Route("/account", func(router chi.Router) {
		//middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)

	})

}
