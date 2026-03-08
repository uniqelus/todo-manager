package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/jackc/pgx/v5/pgxpool"

	httproute "github.com/uniqelus/todo-manager/internal/handlers/http"
	httpserver "github.com/uniqelus/todo-manager/pkg/components/http/server"
	"github.com/uniqelus/todo-manager/pkg/logging"
)

type App struct {
	cfg        *Config
	log        *zap.Logger
	db         *pgxpool.Pool
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

	httpsServer := newHTTPServer(log, &cfg.HTTP)

	db, err := newDatabase(log, &cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("cannot establish connection to database: %w", err)
	}

	return &App{
		cfg:        cfg,
		log:        log,
		db:         db,
		httpServer: httpsServer,
	}, nil
}

func newHTTPServer(
	log *zap.Logger,
	cfg *HTTPConfig,
) *httpserver.Component {
	log.Info("setting http server",
		zap.String("listen_host", cfg.ListenHost),
		zap.String("listen_port", cfg.ListenPort),
		zap.Duration("read_timeout", cfg.ReadTimeout),
		zap.Duration("write_timeout", cfg.WriteTimeout),
		zap.Duration("idle_timeout", cfg.IdleTimeout),
	)

	httpServerAddress := net.JoinHostPort(cfg.ListenHost, cfg.ListenPort)
	httpServer := httpserver.NewComponent(httpServerAddress,
		httpserver.WithHandler(httproute.NewRouter()),
		httpserver.WithReadTimeout(cfg.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.WriteTimeout),
		httpserver.WithIdleTimeout(cfg.IdleTimeout),
	)

	return httpServer
}

func newDatabase(log *zap.Logger, cfg *DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.DBName, cfg.User, cfg.Password,
	)

	pgxCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("parse db config: %w", err)
	}

	pgxCfg.MaxConns = cfg.MaxConns
	pgxCfg.MinConns = cfg.MinConns
	pgxCfg.MaxConnIdleTime = cfg.MaxConnIdle
	pgxCfg.MaxConnLifetime = cfg.MaxConnLifetime

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxCfg)
	if err != nil {
		return nil, fmt.Errorf("create db pool: %w", err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	log.Info("database connected",
		zap.String("host", cfg.Host),
		zap.String("port", cfg.Port),
		zap.String("dbname", cfg.DBName),
	)

	return pool, nil
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

	g.Go(func() error {
		if a.db != nil {
			a.log.Info("closing database pool")
			a.db.Close()
		}
		return nil
	})

	return g.Wait()
}
