package taskrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/uniqelus/todo-manager/internal/domains/server/task"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateTask(ctx context.Context, task *task.Task) error {
	const query = `
		INSERT INTO state.tasks (
			id, title, description, priority, due_date, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	_, err := r.db.Exec(ctx, query,
		task.ID,
		task.Title,
		task.Description,
		task.Priority.String(),
		task.DueDate,
		task.CreatedAt,
		task.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (r *Repository) GetTask(ctx context.Context, id uuid.UUID) (*task.Task, error) {
	const query = `
		SELECT 
			id, title, description, priority,
			due_date, created_at, updated_at
		FROM state.tasks
		WHERE id = $1
	`

	var t task.Task
	var priorityStr string

	err := r.db.QueryRow(ctx, query, id).Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&priorityStr,
		&t.DueDate,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("task not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	t.Priority, err = task.NewPriority(priorityStr)
	if err != nil {
		return nil, fmt.Errorf("invalid priority value in database: %w", err)
	}

	return &t, nil
}
