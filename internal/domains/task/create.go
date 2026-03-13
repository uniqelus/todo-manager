package taskdmn

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
	Priority    Priority
	Recurrence  Recurrence
	DueDate     time.Time
}

func NewCreateTaskOptions(opts ...CreateTaskOption) (*CreateTaskOptions, error) {
	options := &CreateTaskOptions{}
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

func (cto *CreateTaskOptions) validate() error {
	if err := validateTaskTitle(cto.Title); err != nil {
		return err
	}

	if err := validateTaskDescription(cto.Description); err != nil {
		return err
	}

	if err := validateTaskDueDate(cto.DueDate); err != nil {
		return err
	}

	return nil
}

type CreateTaskOption func(*CreateTaskOptions) error

func WithTitle(title string) CreateTaskOption {
	return func(cto *CreateTaskOptions) error {
		cto.Title = title
		return nil
	}
}

func WithDescription(description string) CreateTaskOption {
	return func(cto *CreateTaskOptions) error {
		cto.Description = description
		return nil
	}
}

func WithPriority(priorityValue string) CreateTaskOption {
	return func(cto *CreateTaskOptions) error {
		priority, err := NewPriority(priorityValue)
		if err != nil {
			return err
		}

		cto.Priority = priority
		return nil
	}
}

func WithRecurrence(recurrencePattern string) CreateTaskOption {
	return func(cto *CreateTaskOptions) error {
		recurrence, err := NewRecurrence(recurrencePattern)
		if err != nil {
			return err
		}

		cto.Recurrence = recurrence
		return nil
	}
}

func WithDueDate(dueDate string) CreateTaskOption {
	return func(cto *CreateTaskOptions) error {
		parsedDueDate, err := time.Parse(DueDateLayout, dueDate)
		if err != nil {
			return err
		}

		cto.DueDate = parsedDueDate
		return nil
	}
}
