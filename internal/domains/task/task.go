package taskdmn

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

const (
	MaxTitleLen       int    = 100
	MaxDescriptionLen int    = 1000
	DueDateLayout     string = "2006-01-02"
)

var (
	ErrTitleIsEmpty       = errors.New("title is empty")
	ErrTitleTooLong       = errors.New("title is too long")
	ErrDescriptionTooLong = errors.New("description is too long")
	ErrDueDateInPast      = errors.New("due date in past")
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	State       State      `json:"state"`
	Recurrence  Recurrence `json:"recurrence"`
	Priority    Priority   `json:"priority"`
	DueDate     time.Time  `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   time.Time  `json:"deleted_at"`
}

func validateTaskTitle(title string) error {
	runeCount := utf8.RuneCountInString(title)
	if runeCount == 0 {
		return ErrTitleIsEmpty
	}

	if runeCount > MaxTitleLen {
		return ErrTitleTooLong
	}

	return nil
}

func validateTaskDescription(description string) error {
	if utf8.RuneCountInString(description) > MaxDescriptionLen {
		return ErrDescriptionTooLong
	}

	return nil
}

func validateTaskDueDate(dueDate time.Time) error {
	if dueDate.Before(time.Now()) {
		return ErrDueDateInPast
	}

	return nil
}
