package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mortage-calc/services/calc/internal/app"
	"mortage-calc/services/calc/internal/config"
)

func main() {
	configPath := "../config/config.yaml"
	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config from %s: %+v", configPath, err)
	}

	createCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	app, err := app.New(createCtx, config.Port)
	if err != nil {
		panic(err)
	}

	runCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()
	go func() {
		err := app.Run(context.Background())
		if err != nil {
			log.Fatalf("Error during app run: %+v", err)
		}
	}()
	<-runCtx.Done()

	stopCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err = app.Stop(stopCtx); err != nil {
		panic(err)
	}
}
