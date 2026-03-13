package taskdmn

import (
	"fmt"
	"strings"
)

type Priority string

const (
	UndefinedPriority Priority = "undefined"
	LowPriority       Priority = "low"
	MediumPriority    Priority = "medium"
	HighPriority      Priority = "high"
)

func NewPriority(value string) (Priority, error) {
	var p Priority
	if err := p.Set(value); err != nil {
		return "", err
	}
	return p, nil
}

func (p *Priority) Set(value string) error {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "undefined":
		*p = UndefinedPriority
	case "low":
		*p = LowPriority
	case "medium":
		*p = MediumPriority
	case "high":
		*p = HighPriority
	default:
		return fmt.Errorf("unsupported priority value: %s", value)
	}
	return nil
}

func (p *Priority) String() string {
	return string(*p)
}
