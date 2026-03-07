package httproute

import (
	"github.com/go-chi/chi/v5"

	"github.com/uniqelus/todo-manager/internal/handlers/http/health"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/health", health.NewHandler())

	return router
}
