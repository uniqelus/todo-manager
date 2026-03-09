package task

import (
	"context"
	"time"
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
