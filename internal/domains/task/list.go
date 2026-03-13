package taskdmn

import (
	"context"
	"fmt"
)

const (
	DefaultPageSize = 50
	MinPageSize     = 1
	MaxPageSize     = 1000
)

var (
	ErrPageSizeOutOfRange = fmt.Errorf("page size must be between %d and %d", MinPageSize, MaxPageSize)
)

type Lister interface {
	ListTasks(ctx context.Context, opts *ListTasksOptions) (*ListTaskResult, error)
}

type ListTasksOptions struct {
	PageSize    int
	PageToken   *PageTokenData
	ShowDeleted bool
	Filter      string
	OrderBy     string
}

type ListTaskResult struct {
	Tasks         []*Task
	NextPageToken PageToken
}

func NewListTasksOptions(opts ...ListTasksOption) (*ListTasksOptions, error) {
	options := &ListTasksOptions{PageSize: DefaultPageSize}
	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}
	}

	return options, nil
}

type ListTasksOption func(*ListTasksOptions) error

func WithListPageSize(size int) ListTasksOption {
	return func(lto *ListTasksOptions) error {
		if size < MinPageSize || size > MaxPageSize {
			return ErrPageSizeOutOfRange
		}

		lto.PageSize = size
		return nil
	}
}

func WithListPageToken(token string) ListTasksOption {
	return func(lto *ListTasksOptions) error {
		pageToken, err := PageTokenFromString(token)
		if err != nil {
			return err
		}

		decodedPageToken, err := DecodePageToken(pageToken)
		if err != nil {
			return err
		}

		lto.PageToken = decodedPageToken
		return nil
	}
}

func WithListShowDeleted(show bool) ListTasksOption {
	return func(lto *ListTasksOptions) error {
		lto.ShowDeleted = show
		return nil
	}
}

func WithListFilter(filter string) ListTasksOption {
	return func(lto *ListTasksOptions) error {
		lto.Filter = filter
		return nil
	}
}
