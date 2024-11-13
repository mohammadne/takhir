package cmd

import (
	"context"

	"github.com/mohammadne/takhir/internal/config"
	"github.com/mohammadne/takhir/internal/http"
	"github.com/mohammadne/takhir/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Server struct {
	ports struct {
		monitor int
		client  int
	}

	config *config.Config
	logger *zap.Logger
}

func (server Server) Command(ctx context.Context) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		server.initialize()
		server.run()
		<-ctx.Done()
		server.stop()
	}

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run Takhir web-server",
		Run:   run,
	}

	cmd.Flags().IntVar(&server.ports.monitor, "monitor-port", 8000, "The port the metric and probes are bind to")
	cmd.Flags().IntVar(&server.ports.client, "client-port", 8001, "The server port which handles client requests")

	return cmd
}

func (server *Server) initialize() {
	server.config = config.Load(true)
	server.logger = logger.NewZap(server.config.Logger)

	server.logger.Info("server has been fully initialized")
}

func (server *Server) run() {
	http.New(server.logger).Serve(server.ports.monitor, server.ports.client)
}

func (server *Server) stop() {
	server.logger.Warn("interruption signal recieved, gracefully shutdown the server")
}
