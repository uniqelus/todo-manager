package logging

import (
	"fmt"

	"go.uber.org/zap"
)

func MustLogger(opts ...Option) *zap.Logger {
	log, err := NewLogger(opts...)
	if err != nil {
		log = zap.NewNop()
	}

	return log
}

func NewLogger(opts ...Option) (*zap.Logger, error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	if options.env == "nop" {
		return zap.NewNop(), nil
	}

	var loggerCfg zap.Config
	switch options.env {
	case "prod":
		loggerCfg = zap.NewProductionConfig()
	case "dev":
		loggerCfg = zap.NewDevelopmentConfig()
	}

	level, err := zap.ParseAtomicLevel(options.level)
	if err != nil {
		return nil, fmt.Errorf("cannot parse atomic level: %w", err)
	}

	loggerCfg.Level = level
	loggerCfg.Encoding = options.encoding
	loggerCfg.OutputPaths = options.paths

	return loggerCfg.Build()
}
