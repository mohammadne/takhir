package http

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/mohammadne/takhir/internal/api/http/handlers"
	"github.com/mohammadne/takhir/internal/api/http/i18n"
	"github.com/mohammadne/takhir/internal/usecases"
)

type Server struct {
	logger *zap.Logger

	monitorApp *fiber.App
	requestApp *fiber.App
}

func New(log *zap.Logger, i18n i18n.I18N,
	categoriesUsecase usecases.Categories,
) *Server {
	server := &Server{logger: log}
	fiberConfig := fiber.Config{DisableStartupMessage: true}

	// ----------------------------------------- Monitor Endpoints

	server.monitorApp = fiber.New(fiberConfig)

	server.monitorApp.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	handlers.NewHealthz(server.monitorApp, log)

	// ----------------------------------------- Request Endpoints

	server.requestApp = fiber.New(fiberConfig)

	v1 := server.requestApp.Group("api/v1")
	handlers.NewCategories(v1, log, i18n, categoriesUsecase)

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

	for _, r := range runnables {
		go func() {
			address := fmt.Sprintf("0.0.0.0:%d", r.port)
			fields := []zapcore.Field{
				zap.String("address", address),
				zap.String("description", r.description),
			}

			server.logger.Info("starting server", fields...)
			err := r.app.Listen(address)
			fields = append(fields, zap.Error(err))
			server.logger.Fatal("error resolving server", fields...)
		}()
	}

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, r := range runnables {
		if err := r.app.ShutdownWithContext(shutdownCtx); err != nil {
			server.logger.Error("error shutdown http server", zap.Error(err))
		}
	}

	server.logger.Warn("gracefully shutdown the https servers")
}
