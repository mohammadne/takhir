package cmd

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/mohammadne/takhir/internal/config"
	"github.com/mohammadne/takhir/internal/http"
	"github.com/mohammadne/takhir/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Server struct {
	ctx context.Context

	ports struct {
		master int
	}

	config *config.Config
	logger *zap.Logger
}

func (server Server) Command() *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		var stop context.CancelFunc
		server.ctx, stop = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		server.initialize()
		server.start()
		defer server.stop()

		<-server.ctx.Done()
		server.logger.Warn("Got interruption signal, gracefully shutdown the server")
	}

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run Takhir web-server",
		Run:   run,
	}

	cmd.Flags().IntVar(&server.ports.master, "master-port", 8000, "The port the metric and probe endpoints binds to")

	return cmd
}

func (server *Server) initialize() {
	server.config = config.Load(true)
	server.logger = logger.NewZap(server.config.Logger)

	server.logger.Info("initialize")
}

func (server *Server) start() {
	server.logger.Info("start")

	http.New(server.logger).
		Serve(server.ports.master, 8080)
}

func (server *Server) stop() {
	server.logger.Info("stop")
}
