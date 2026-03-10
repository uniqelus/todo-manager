package task

import (
	"errors"
	"fmt"
)

type Priority string

const (
	UndefinedPriority Priority = "undefined"
	LowPriority       Priority = "low"
	MediumPriority    Priority = "medium"
	HighPriority      Priority = "high"
)

func NewPriority(value string) (Priority, error) {
	value = normalizeValue(value)
	if err := validatePriority(value); err != nil {
		return "", fmt.Errorf("cannot create priority instance with expected value: %w", err)
	}

	return Priority(value), nil
}

func (p *Priority) Set(value string) error {
	value = normalizeValue(value)
	if err := validatePriority(value); err != nil {
		return fmt.Errorf("cannot set expected value as priority: %w", err)
	}
	*p = Priority(value)

	return nil
}

func (p *Priority) String() string {
	return string(*p)
}

func normalizeValue(value string) string {
	if value == "" {
		return string(UndefinedPriority)
	}

	return value
}

func validatePriority(value string) error {
	switch value {
	case "undefined", "low", "medium", "high":
		return nil
	default:
		return errors.New("unsupported priotity value")
	}
}
