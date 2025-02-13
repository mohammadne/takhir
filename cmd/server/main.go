package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"sync"
	"syscall"

	"github.com/mohammadne/takhir/cmd"
	"github.com/mohammadne/takhir/internal/api/http"
	"github.com/mohammadne/takhir/internal/config"
	"github.com/mohammadne/takhir/internal/core"
	"github.com/mohammadne/takhir/pkg/databases/postgres"
	"github.com/mohammadne/takhir/pkg/databases/redis"
	"github.com/mohammadne/takhir/pkg/observability/logger"
	"go.uber.org/zap"
)

func main() {
	monitorPort := flag.Int("monitor-port", 8087, "The server port which handles monitoring endpoints (default: 8087)")
	requestPort := flag.Int("request-port", 8088, "The server port which handles http requests (default: 8088)")
	environmentRaw := flag.String("environment", "", "The environment (default: local)")
	flag.Parse() // Parse the command-line flags

	var cfg config.Config
	var err error

	switch core.ToEnvironment(*environmentRaw) {
	case core.EnvironmentLocal:
		cfg, err = config.LoadDefaults(true)
	default:
		cfg, err = config.Load(true)
	}

	if err != nil {
		log.Fatalf("failed to load config: \n%v", err)
	}

	logger, err := logger.New(cfg.Logger)
	if err != nil {
		log.Fatalf("failed to initialize logger: \n%v", err)
	}

	logger.Info("BuildInfo", zap.Any("data", cmd.BuildInfo()))

	postgres, err := postgres.Open(cfg.Postgres, core.Namespace, core.System)
	if err != nil {
		logger.Fatal(`error connecting to postgres database`, zap.Error(err))
	}
	_ = postgres

	redis, err := redis.Open(cfg.Redis)
	if err != nil {
		logger.Fatal(`error connecting to redis database`, zap.Error(err))
	}
	_ = redis

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	var wg sync.WaitGroup

	wg.Add(1)
	go http.New(logger).Serve(ctx, &wg, *monitorPort, *requestPort)

	<-ctx.Done()
	wg.Wait()
	logger.Warn("interruption signal recieved, gracefully shutdown the server")
}
