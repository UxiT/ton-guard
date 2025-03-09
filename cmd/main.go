package main

import (
	"decard/config"
	"decard/config/container"
	"decard/internal/infrastructure/bootstrap"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.Cfg

	registry := container.NewContainer(cfg)

	defer registry.DB.Close()

	application := bootstrap.New(registry, cfg.ServerAddress)

	go application.HTTPSrv.MusRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	systemCall := <-stop

	registry.Logger.Info("stopping application", slog.String("signal", systemCall.String()))

	application.HTTPSrv.Stop()

	registry.Logger.Info("application stopped")
}
