package libconfig

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

func ReadFromFile[T any](path string) (*T, error) {
	var cfg T
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("cannot read configuration from file: %w", err)
	}
	return &cfg, nil
}

func ReadFromEnv[T any]() (*T, error) {
	var cfg T
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("cannot read configuration from env: %w", err)
	}
	return &cfg, nil
}
