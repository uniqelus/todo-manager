package task

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Service interface {
}

func NewRouter(log *zap.Logger, service Service) *chi.Mux {
	router := chi.NewRouter()
	return router
}
