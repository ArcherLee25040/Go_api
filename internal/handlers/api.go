package handlers

import "github.com/go-admin-team/go-admin-core/sdk/middleware"

// "github.com/go-chi/chi"
// chimiddle "github.com/go-chi/chi/middleware"
// "github.com/ArcherLee25040/go_api/internal/middleware"

func Handler(r *chi.Mux) {
	//Global middleware
	r.Use(chimiddle.StriptSlashes)

	r.Route("/account", func(router chi.Router) {
		//Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
