package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/hemanG-G/GO_MINI_PROJECTS/tree/main/BasicAPI/internal/middleware"
)

// the handler function that accepts router
func Handler(r *chi.Mux) { // "H" capital , for public function

	//Global middleware
	r.Use(chimiddle.StripSlashes) // strip slashes at end of URL

	// creating /account route
	r.Route("/account", func(router chi.Router) { // creating a subroute /coins
		//middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance) // "GET" route

	})

}
