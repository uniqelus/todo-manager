package task

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"       validate:"required,min=1,max=255"`
	Description string    `json:"description" validate:"max=1000"`
	Priority    Priority  `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewTask creates new task instance with gotten title, decription, dueDate and priotiry.
func NewTask(title string, description string, dueDate time.Time, priority string) (*Task, error) {
	typedPriority, err := NewPriority(priority)
	if err != nil {
		return nil, err
	}

	task := &Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Priority:    typedPriority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	validator := validator.New()
	if err = validator.Struct(task); err != nil {
		return nil, fmt.Errorf("cannot create task: %w", err)
	}

	return task, nil
}
