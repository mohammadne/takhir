package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/mohammadne/takhir/cmd"
	"github.com/spf13/cobra"
)

func main() {
	const description = "Takhir main entrypoint"
	root := &cobra.Command{Short: description}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	root.AddCommand(
		cmd.Server{}.Command(ctx),
		cmd.Migration{}.Command(ctx),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute root command: \n%v", err)
	}
}
