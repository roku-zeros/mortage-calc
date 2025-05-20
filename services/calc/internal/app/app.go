package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/roku-zeros/mortage-calc/lib/middleware"
	"github.com/roku-zeros/mortage-calc/services/calc/internal/providers"
	storage "github.com/roku-zeros/mortage-calc/services/calc/internal/repository/database"
	"github.com/roku-zeros/mortage-calc/services/calc/internal/server"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

type App struct {
	server Server
}

func New(ctx context.Context, port string) (*App, error) {
	storage := storage.NewStorage(ctx)
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
