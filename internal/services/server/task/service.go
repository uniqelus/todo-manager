package task

import (
	"context"
	"fmt"

	"github.com/uniqelus/todo-manager/internal/domain/server/task"
)

//go:generate mockery --name=Repository --output=mocks --filename=repository.go
type Repository interface {
	CreateTask(ctx context.Context, task *task.Task) error
}

type Service struct {
	taskRepository Repository
}

func NewService(tr Repository) *Service {
	return &Service{
		taskRepository: tr,
	}
}

func (s *Service) CreateTask(ctx context.Context, options *task.CreateTaskOptions) (*task.Task, error) {
	task, err := task.NewTask(options.Title, options.Description, options.DueDate, options.Priority)
	if err != nil {
		return nil, fmt.Errorf("cannot create task model: %w", err)
	}

	if err = s.taskRepository.CreateTask(ctx, task); err != nil {
		return nil, fmt.Errorf("cannot create task: %w", err)
	}

	return task, nil
}
