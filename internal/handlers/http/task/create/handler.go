package create

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"go.uber.org/zap"

	"github.com/uniqelus/todo-manager/internal/domains/server/task"
	"github.com/uniqelus/todo-manager/internal/handlers/http/helpers"
)

type Request struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"due_date"`
}

type Response struct {
	Task  task.Task `json:"task,omitzero"`
	Error string    `json:"error,omitempty"`
}

//go:generate mockery --name=TaskCreator --output=mocks --filename=task_creator.go
type TaskCreator interface {
	task.Creator
}

var ErrFailedToCreateTask = errors.New("failed to create task")

func NewHandler(log *zap.Logger, tc TaskCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hLog := log.With(zap.String("operation", "task.create"))

		var req Request
		if err := helpers.DecodeRequest(r, &req); err != nil {
			hLog.Error(err.Error())

			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, Response{Error: err.Error()})

			return
		}

		hLog.Info("request body decoded",
			zap.String("title", req.Title),
			zap.String("description", req.Description),
			zap.String("priority", req.Priority),
			zap.Time("due_date", req.DueDate),
		)

		options := &task.CreateTaskOptions{
			Title:       req.Title,
			Description: req.Description,
			Priority:    req.Priority,
			DueDate:     req.DueDate,
		}

		created, err := tc.CreateTask(r.Context(), options)
		if err != nil {
			hLog.Error(ErrFailedToCreateTask.Error(), zap.Error(err))

			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, Response{Error: ErrFailedToCreateTask.Error()})

			return
		}

		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, Response{Task: *created})
	}
}
