package logging

import "strings"

type options struct {
	env      string
	level    string
	encoding string
	paths    []string
}

func defaultOptions() *options {
	return &options{
		env:      "prod",
		level:    "info",
		encoding: "json",
		paths:    []string{"stdout"},
	}
}

type Option func(*options)

func WithEnv(env string) Option {
	return func(o *options) {
		o.env = strings.ToLower(env)
	}
}

func WithLevel(level string) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithEncoding(encoding string) Option {
	return func(o *options) {
		o.encoding = encoding
	}
}

func WithPaths(paths ...string) Option {
	return func(o *options) {
		o.paths = append(o.paths, paths...)
	}
}
