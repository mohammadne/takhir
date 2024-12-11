package cmd

import (
	"context"

	"github.com/mohammadne/takhir/internal/config"
	"github.com/mohammadne/takhir/internal/http"
	"github.com/mohammadne/takhir/pkg/logger"
	"github.com/mohammadne/takhir/pkg/postgres"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Server struct {
	ports struct {
		monitor int
		client  int
	}

	config   *config.Config
	logger   *zap.Logger
	postgres *postgres.Postgres
}

func (s Server) Command(ctx context.Context) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		s.initialize()
		s.run()
		<-ctx.Done()
		s.stop()
	}

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run Takhir web-server",
		Run:   run,
	}

	cmd.Flags().IntVar(&s.ports.monitor, "monitor-port", 8000, "The port the metric and probes are bind to")
	cmd.Flags().IntVar(&s.ports.client, "client-port", 8001, "The server port which handles client requests")

	return cmd
}

func (s *Server) initialize() {
	s.config = config.Load(true)
	s.logger = logger.NewZap(s.config.Logger)

	postgres, err := postgres.Open(s.config.Postgres, "file://hacks/migrations")
	if err != nil {
		s.logger.Fatal("error connecting to postgresql database", zap.Error(err))
	}
	s.postgres = postgres

	s.logger.Info("server has been fully initialized")
}

func (s *Server) run() {
	http.New(s.logger).Serve(s.ports.monitor, s.ports.client)
}

func (s *Server) stop() {
	s.logger.Warn("interruption signal recieved, gracefully shutdown the server")
}
