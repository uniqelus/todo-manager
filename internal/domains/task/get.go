package taskdmn

import (
	"context"

	"github.com/google/uuid"
)

type Getter interface {
	GetTask(ctx context.Context, opts *GetTaskOptions) (*Task, error)
}

type GetTaskOptions struct {
	ID          uuid.UUID
	ShowDeleted bool
}

func NewGetTaskOptions(opts ...GetTaskOption) (*GetTaskOptions, error) {
	options := &GetTaskOptions{}
	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}
	}

	return options, nil
}

type GetTaskOption func(*GetTaskOptions) error

func WithGetTaskByID(id string) GetTaskOption {
	return func(gto *GetTaskOptions) error {
		taskID, err := uuid.Parse(id)
		if err != nil {
			return err
		}

		gto.ID = taskID
		return nil
	}
}

func WithGetDeleted(show bool) GetTaskOption {
	return func(gto *GetTaskOptions) error {
		gto.ShowDeleted = show
		return nil
	}
}
