package task

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/uniqelus/todo-manager/internal/handlers/http/task/create"
)

type Service interface {
	create.TaskCreator
}

func NewRouter(log *zap.Logger, service Service) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/task", create.NewHandler(log, service))

	return router
}
