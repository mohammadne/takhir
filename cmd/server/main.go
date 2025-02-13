package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mohammadne/takhir/internal/api/http"
	"github.com/mohammadne/takhir/internal/config"
	"github.com/mohammadne/takhir/internal/core"
	"github.com/mohammadne/takhir/pkg/databases/postgres"
	"github.com/mohammadne/takhir/pkg/observability/logger"
)

func main() {
	monitorPort := flag.Int("monitor-port", 8087, "The server port which handles monitoring endpoints (default: 8087)")
	requestPort := flag.Int("request-port", 8088, "The server port which handles http requests (default: 8088)")
	environmentRaw := flag.String("environment", "", "The environment (default: local)")
	flag.Parse() // Parse the command-line flags

	environment := core.ToEnvironment(*environmentRaw)

	var cfg config.Config
	var err error

	switch environment {
	case core.EnvironmentLocal:
		cfg, err = config.LoadDefaults(true)
		if err != nil {
			panic(err)
		}
	default:
		cfg, err = config.Load(true)
		if err != nil {
			panic(err)
		}
	}

	logger, err := logger.New(cfg.Logger)
	if err != nil {
		panic(err)
	}

	postgres, err := postgres.Open(cfg.Postgres, core.Namespace, core.System)
	if err != nil {
		slog.Error(`error connecting to postgres database`, `Err`, err)
		os.Exit(1)
	}
	_ = postgres

	// redis, err := redis.Open(cfg.Redis)
	// if err != nil {
	// 	slog.Error(`error connecting to redis database`, `Err`, err)
	// 	os.Exit(1)
	// }

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	var wg sync.WaitGroup

	wg.Add(1)
	go http.New(logger).Serve(ctx, &wg, *monitorPort, *requestPort)

	<-ctx.Done()
	wg.Wait()
	slog.Warn("interruption signal recieved, gracefully shutdown the server")
}
