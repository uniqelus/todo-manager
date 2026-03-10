package task

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Creator interface {
	CreateTask(ctx context.Context, opts *CreateTaskOptions) (*Task, error)
}

type CreateTaskOptions struct {
	Title       string
	Description string
	Priority    string
	DueDate     time.Time
}

type Getter interface {
	GetTask(ctx context.Context, opts *GetTaskOpotions) (*Task, error)
}

type GetTaskOpotions struct {
	ID uuid.UUID
}
