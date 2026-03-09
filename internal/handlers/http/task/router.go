package task

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/uniqelus/todo-manager/internal/handlers/http/task/create"
	"github.com/uniqelus/todo-manager/internal/handlers/http/task/get"
)

type Service interface {
	create.TaskCreator
	get.TaskGetter
}

func NewRouter(log *zap.Logger, service Service) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/task", create.NewHandler(log, service))
	router.Get("/task", get.NewHandler(log, service))

	return router
}
