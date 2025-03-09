package bootstrap

import (
	"decard/config/container"
	"decard/internal/infrastructure/bootstrap/http"
)

type App struct {
	HTTPSrv *http.App
}

func New(
	container *container.Container,
	address string,
) *App {
	httpSrv := http.New(container.Logger, address, container.Router)

	return &App{
		HTTPSrv: httpSrv,
	}
}
