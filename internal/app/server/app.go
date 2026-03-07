package server

import (
	"context"
	"errors"
	"net"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	httproute "github.com/uniqelus/todo-manager/internal/handlers/http"
	httpserver "github.com/uniqelus/todo-manager/pkg/components/http/server"
	"github.com/uniqelus/todo-manager/pkg/logging"
)

type App struct {
	log *zap.Logger

	httpServer *httpserver.Component
}

func NewApp(cfg *Config) (*App, error) {
	if cfg == nil {
		return nil, errors.New("configuration is required")
	}

	log := logging.MustLogger(
		logging.WithEnv(cfg.Logging.Env),
		logging.WithLevel(cfg.Logging.Level),
		logging.WithEncoding(cfg.Logging.Encoding),
		logging.WithPaths(cfg.Logging.Paths...),
	).With(zap.String("service", "todo-manager-server"))

	log.Info("setting http server",
		zap.String("listen_host", cfg.HTTP.ListenHost),
		zap.String("listen_port", cfg.HTTP.ListenPort),
		zap.Duration("read_timeout", cfg.HTTP.ReadTimeout),
		zap.Duration("write_timeout", cfg.HTTP.WriteTimeout),
		zap.Duration("idle_timeout", cfg.HTTP.IdleTimeout),
	)

	httpServerAddress := net.JoinHostPort(cfg.HTTP.ListenHost, cfg.HTTP.ListenPort)
	httpServer := httpserver.NewComponent(httpServerAddress,
		httpserver.WithHandler(httproute.NewRouter()),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithIdleTimeout(cfg.HTTP.IdleTimeout),
	)

	return &App{
		log:        log,
		httpServer: httpServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		a.log.Info("running http server")
		return a.httpServer.Run(gCtx)
	})

	return g.Wait()
}

func (a *App) Stop(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		a.log.Info("stopping http server")
		return a.httpServer.Stop(gCtx)
	})

	return g.Wait()
}
