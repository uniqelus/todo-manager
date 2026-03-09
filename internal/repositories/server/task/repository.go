package task

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/uniqelus/todo-manager/internal/domain/server/task"
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
