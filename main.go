package main

import (
	"log"

	"github.com/mohammadne/takhir/cmd"
	"github.com/spf13/cobra"
)

func main() {
	const description = "Takhir main entrypoint"
	root := &cobra.Command{Short: description}

	root.AddCommand(
		cmd.Server{}.Command(),
		cmd.Migration{}.Command(),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute root command: \n%v", err)
	}
}
