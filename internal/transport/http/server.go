package http

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/mohammadne/takhir/internal/transport/http/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Server struct {
	logger *zap.Logger

	monitorApp *fiber.App
	requestApp *fiber.App
}

func New(log *zap.Logger) *Server {
	server := &Server{logger: log}
	fiberConfig := fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal}

	// ----------------------------------------- Monitor Endpoints

	server.monitorApp = fiber.New(fiberConfig)

	server.monitorApp.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	handlers.NewHealthz(server.monitorApp, log)

	// ----------------------------------------- Request Endpoints

	server.requestApp = fiber.New(fiberConfig)

	v1 := server.requestApp.Group("api/v1")
	handlers.NewCategories(v1, log)
	handlers.NewItems(v1, log)

	return server
}

func (server *Server) Serve(ctx context.Context, wg *sync.WaitGroup, monitor, request int) {
	runnables := []struct {
		port        int
		app         *fiber.App
		description string
	}{
		{monitor, server.monitorApp, "monitor server"},
		{request, server.requestApp, "request server"},
	}

	for _, runnable := range runnables {
		go func() {
			address := fmt.Sprintf("0.0.0.0:%d", runnable.port)
			fields := []zapcore.Field{zap.String("address", address),
				zap.String("description", runnable.description)}

			server.logger.Info("starting server", fields...)
			err := runnable.app.Listen(address)
			fields = append(fields, zap.Error(err))
			server.logger.Fatal("error resolving server", fields...)
		}()
	}

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, runnable := range runnables {
		if err := runnable.app.ShutdownWithContext(shutdownCtx); err != nil {
			server.logger.Error("error shutdown http server", zap.Error(err))
		}
	}

	server.logger.Warn("gracefully shutdown the https servers")
}
