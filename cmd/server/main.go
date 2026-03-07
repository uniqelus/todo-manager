package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/uniqelus/todo-manager/internal/app/server"
	libconfig "github.com/uniqelus/todo-manager/pkg/config"
	liberrors "github.com/uniqelus/todo-manager/pkg/errors"
)

func main() {
	var path string

	flag.StringVar(&path, "config", "", "path to configuration file")
	flag.Parse()

	cfg := liberrors.Must(readConfig(path))

	app := liberrors.Must(server.NewApp(cfg))

	runCtx, runCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer runCancel()

	liberrors.Try(app.Run(runCtx))
	<-runCtx.Done()

	runCancel()

	stopCtx, stopCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopCancel()

	liberrors.Try(app.Stop(stopCtx))
}

func readConfig(path string) (*server.Config, error) {
	if path != "" {
		return libconfig.ReadFromFile[server.Config](path)
	}
	return libconfig.ReadFromEnv[server.Config]()
}
