package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type App struct {
	logger     *slog.Logger
	httpServer *http.Server
}

func New(
	logger *slog.Logger,
	address string,
	router *mux.Router,
) *App {
	return &App{
		logger:     logger,
		httpServer: &http.Server{Addr: address, Handler: router},
	}
}

func (a *App) MusRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "app.http.run"

	logger := a.logger.With(
		slog.String("operation", op),
		slog.String("address", a.httpServer.Addr),
	)

	logger.Info("starting http server")

	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "app.http.stop"

	a.logger.Info("stopping http server", slog.String("operation", op))

	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}
