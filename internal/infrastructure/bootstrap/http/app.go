package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
)

type App struct {
	logger     *zerolog.Logger
	httpServer *http.Server
}

func New(
	logger *zerolog.Logger,
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

	a.logger.Info().Str("operation", op).Str("address", a.httpServer.Addr).Msg("starting http server")

	if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "app.http.stop"

	a.logger.Info().Str("operation", op).Msg("stopping http server")

	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}
