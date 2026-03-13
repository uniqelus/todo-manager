package taskdmn

import (
	"context"

	"github.com/google/uuid"
)

type Deleter interface {
	DeleteTask(ctx context.Context, opts *DeleteTaskOptions) (*Task, error)
}

type DeleteTaskOptions struct {
	ID    uuid.UUID
	Force bool
}

func NewDeleteTaskOptions(opts ...DeleteTaskOption) (*DeleteTaskOptions, error) {
	options := &DeleteTaskOptions{}
	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}
	}

	return options, nil
}

type DeleteTaskOption func(*DeleteTaskOptions) error

func WithDeleteTaskByID(id string) DeleteTaskOption {
	return func(dto *DeleteTaskOptions) error {
		taskID, err := uuid.Parse(id)
		if err != nil {
			return err
		}

		dto.ID = taskID
		return nil
	}
}

func WithForceDelete(force bool) DeleteTaskOption {
	return func(dto *DeleteTaskOptions) error {
		dto.Force = force
		return nil
	}
}
