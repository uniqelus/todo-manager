package httproute

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/uniqelus/todo-manager/internal/handlers/http/health"
	taskroutes "github.com/uniqelus/todo-manager/internal/handlers/http/task"
)

func NewRouter(
	log *zap.Logger,
	taskService taskroutes.Service,
) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/health", health.NewHandler())

	router.Mount("/api", taskroutes.NewRouter(log, taskService))

	return router
}
