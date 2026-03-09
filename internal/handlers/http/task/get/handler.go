package get

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/uniqelus/todo-manager/internal/domain/server/task"
)

type Response struct {
	Task  task.Task `json:"task,omitzero"`
	Error string    `json:"error,omitempty"`
}

type TaskGetter interface {
	task.Getter
}

func NewHandler(log *zap.Logger, tg TaskGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hLog := log.With(zap.String("operation", "task.get"))

		taskID, err := uuid.Parse(r.URL.Query().Get("id"))
		if err != nil {
			hLog.Error("invalid task id format", zap.Error(err))

			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, Response{Error: "invalid task id format"})

			return
		}

		opts := &task.GetTaskOpotions{
			ID: taskID,
		}

		retrievedTask, err := tg.GetTask(r.Context(), opts)
		if err != nil {
			if errors.Is(err, task.ErrTaskNotFound) {
				hLog.Error("cannot found task", zap.Error(err))

				w.WriteHeader(http.StatusNotFound)
				render.JSON(w, r, Response{Error: "not found"})

				return
			}

			hLog.Error("cannot retrive task", zap.Error(err))

			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, Response{Error: "cannot retrieve task"})

			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, Response{Task: *retrievedTask})
	}
}
