package http

import (
	"encoding/json"
	"fmt"

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
	clientApp  *fiber.App
}

func New(log *zap.Logger) *Server {
	server := &Server{logger: log}
	fiberConfig := fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal}

	// ----------------------------------------- Monitor Endpoints

	server.monitorApp = fiber.New(fiberConfig)

	server.monitorApp.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	handlers.NewHealthz(server.monitorApp, log)

	// ----------------------------------------- Client Endpoints

	server.clientApp = fiber.New(fiberConfig)

	v1 := server.clientApp.Group("api/v1")
	handlers.NewCategories(v1, log)
	handlers.NewItems(v1, log)

	return server
}

func (server *Server) Serve(monitor, client int) {
	runnables := []struct {
		port        int
		app         *fiber.App
		description string
	}{
		{monitor, server.monitorApp, "monitor server"},
		{client, server.clientApp, "client server"},
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
}
