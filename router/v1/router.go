package v1

import "github.com/go-chi/chi/v5"

func RegisterRoutes(router *chi.Mux) {
	router.Get("/", getHomeApi)
}
