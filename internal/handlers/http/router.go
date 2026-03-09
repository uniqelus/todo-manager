package httproute

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/uniqelus/todo-manager/internal/handlers/http/health"
	mw "github.com/uniqelus/todo-manager/internal/handlers/http/middleware"
	taskroutes "github.com/uniqelus/todo-manager/internal/handlers/http/task"
)

func NewRouter(
	log *zap.Logger,
	taskService taskroutes.Service,
) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		middleware.Recoverer,
		mw.Logging(log),
	)

	router.Get("/health", health.NewHandler())

	router.Mount("/api", taskroutes.NewRouter(log, taskService))

	return router
}
