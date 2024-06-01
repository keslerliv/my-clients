package routes

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/keslerliv/my-clients/internal/handlers"
)

func LoadRoutes() http.Handler {
	router := chi.NewRouter()

	// index route
	router.Get("/", handlers.HomeGet)

	// client routes
	router.Route("/client", func(r chi.Router) {
		r.Get("/", handlers.ClientList)
		r.Get("/{id}", handlers.ClientGet)
		r.Post("/", handlers.ClientCreate)
		r.Put("/{id}", handlers.ClientUpdate)
		r.Delete("/{id}", handlers.ClientDelete)
		r.Post("/upload", handlers.CreateClientsFromTXT)
	})

	return router
}
