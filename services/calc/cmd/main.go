package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/roku-zeros/mortage-calc/services/calc/internal/app"
	"github.com/roku-zeros/mortage-calc/services/calc/internal/config"
)

func main() {
	configPath := flag.String("config", "/config.yaml", "Path to the config file")
	flag.Parse()

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config from %s: %+v", *configPath, err)
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
