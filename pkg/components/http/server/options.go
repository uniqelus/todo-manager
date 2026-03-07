package server

import (
	"net/http"
	"time"
)

type options struct {
	handler      http.Handler
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
}

func defaultOptions() *options {
	return &options{}
}

type Option func(*options)

func WithHandler(handler http.Handler) Option {
	return func(o *options) {
		o.handler = handler
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.readTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.writeTimeout = timeout
	}
}

func WithIdleTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.idleTimeout = timeout
	}
}
