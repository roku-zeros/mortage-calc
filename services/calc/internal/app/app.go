package app

import (
	"context"
	"fmt"
	"mortage-calc/lib/middleware"
	"mortage-calc/services/calc/internal/providers"
	storage "mortage-calc/services/calc/internal/repository/database"
	"mortage-calc/services/calc/internal/server"
	"net/http"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

type App struct {
	server Server
}

func New(ctx context.Context, port string) (*App, error) {
	storage, _ := storage.NewStorage(ctx)
	taskProvider := providers.NewMortageProvider(storage)

	server := server.New(taskProvider)
	mux := http.NewServeMux()
	server.RegisterRoutes(mux)

	httpServer := &http.Server{
		Addr:    port,
		Handler: middleware.RequestLogger(mux),
	}

	return &App{
		server: httpServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	fmt.Println("App started")
	return a.server.ListenAndServe()
}

func (a *App) Stop(ctx context.Context) error {
	fmt.Println("App stoped")
	return a.server.Shutdown(ctx)
}
