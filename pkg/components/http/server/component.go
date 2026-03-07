package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type Component struct {
	address string
	server  *http.Server
}

func NewComponent(address string, opts ...Option) *Component {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	return &Component{
		address: address,
		server: &http.Server{
			Handler:      options.handler,
			ReadTimeout:  options.readTimeout,
			WriteTimeout: options.writeTimeout,
			IdleTimeout:  options.idleTimeout,
		},
	}
}

func (c *Component) Run(ctx context.Context) error {
	errCh := make(chan error)
	go func() {
		lis, err := net.Listen("tcp", c.address) //nolint:noctx // this is why
		if err != nil {
			errCh <- fmt.Errorf("cannot listen expected address '%s': %w", c.address, err)
			return
		}

		select {
		case errCh <- c.server.Serve(lis):
		case <-ctx.Done():
		}

		close(errCh)
	}()

	select {
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	case <-ctx.Done():
		return nil
	}
}

func (c *Component) Stop(ctx context.Context) error {
	return c.server.Shutdown(ctx)
}
