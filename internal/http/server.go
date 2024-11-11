package http

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Server struct {
	logger *zap.Logger
	// repository repository.Repository

	masterApp *fiber.App
	clientApp *fiber.App
}

func New(log *zap.Logger) *Server {
	server := &Server{logger: log}
	fiberConfig := fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal}

	// ----------------------------------------- Master Endpoints

	server.masterApp = fiber.New(fiberConfig)

	server.masterApp.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	server.masterApp.Get("/healthz/liveness", server.liveness)
	server.masterApp.Get("/healthz/readiness", server.readiness)

	// ----------------------------------------- Client Endpoints

	server.clientApp = fiber.New(fiberConfig)

	v1 := server.clientApp.Group("api/v1")
	_ = v1

	// auth := v1.Group("auth")
	// auth.Post("/register", server.register)
	// auth.Post("/login", server.login)

	// contacts := v1.Group("contacts", server.fetchUserId)
	// contacts.Get("/", server.getContacts)
	// contacts.Post("/", server.createContact)
	// contacts.Get("/:id", server.getContact)
	// contacts.Put("/:id", server.updateContact)
	// contacts.Delete("/:id", server.deleteContact)

	return server
}

func (server *Server) Serve(master, client int) {
	runnables := []struct {
		port        int
		app         *fiber.App
		description string
	}{
		{master, server.masterApp, "master server"},
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
