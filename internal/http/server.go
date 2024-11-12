package http

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/mohammadne/takhir/pkg/stackerr"
)

type Server struct {
	logger *zap.Logger
	// repository repository.Repository

	masterApp *fiber.App
	clientApp *fiber.App
}

func SampleError() error {
	// pc := make([]uintptr, 15)
	// n := runtime.Callers(2, pc)
	// frames := runtime.CallersFrames(pc[:n])
	// frame, _ := frames.Next()

	// trace := fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)

	// err := errors.New("some error happened")
	return stackerr.Wrap(child1(), "error in SampleError")

	// return tracerr.Wrap(err)
	// return fmt.Errorf("child module error at someFunction: %w", err)
	// return fmt.Errorf("child module error at someFunction: %w", err)
}

func child1() error {
	return stackerr.Wrap(child2(), "error while calling child3")
}

func child2() error {
	return stackerr.Wrap(child3(), "error while calling child3")
}

func child3() error {
	return errors.New("error from 3rd module")
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
