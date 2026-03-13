package taskdmn

import (
	"errors"
	"strings"
)

type State string

const (
	UndefinedTaskState  State = "undefined"
	ToDoTaskState       State = "todo"
	InProgressTaskState State = "in progress"
	DoneTaskState       State = "done"
	CancelledTaskState  State = "canceled"
	DeletedTaskState    State = "deleted"
)

func NewTaskState(value string) (State, error) {
	var s State
	if err := s.Set(value); err != nil {
		return "", err
	}
	return s, nil
}

func (s *State) Set(value string) error {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "undefined":
		*s = UndefinedTaskState
	case "todo":
		*s = ToDoTaskState
	case "in progress":
		*s = InProgressTaskState
	case "done":
		*s = DoneTaskState
	case "canceled":
		*s = CancelledTaskState
	case "deleted":
		*s = DeletedTaskState
	default:
		return errors.New("invalid task state")
	}
	return nil
}

func (s *State) String() string {
	return string(*s)
}
