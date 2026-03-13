package taskdmn

import (
	"context"
	"errors"
)

type Updater interface {
	UpdateTask(ctx context.Context, opts *UpdateTaskOptions) (*Task, error)
}

type UpdateTaskOptions struct {
	Data  *Task
	Paths []string
}

type UpdateTaskOption func(*UpdateTaskOptions) error

func NewUpdateTaskOptions(opts ...UpdateTaskOption) (*UpdateTaskOptions, error) {
	options := &UpdateTaskOptions{}
	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}
	}

	if err := options.validate(); err != nil {
		return nil, err
	}

	return options, nil
}

func (uto *UpdateTaskOptions) validate() error {
	if uto.Data == nil {
		return errors.New("update task data is required")
	}

	if len(uto.Paths) == 0 {
		return errors.New("update paths are required")
	}

	return nil
}

func WithUpdateData(data *Task) UpdateTaskOption {
	return func(uto *UpdateTaskOptions) error {
		uto.Data = data
		return nil
	}
}

func WithUpdatePaths(paths []string) UpdateTaskOption {
	return func(uto *UpdateTaskOptions) error {
		uto.Paths = paths
		return nil
	}
}
