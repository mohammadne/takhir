package cmd

import (
	"github.com/spf13/cobra"
)

type Migration struct{}

func (m Migration) Command() *cobra.Command {
	run := func(_ *cobra.Command, args []string) {
		// m.main(config.Load(true), args, trap)
	}

	return &cobra.Command{
		Use:       "migration",
		Short:     "run migrations against database",
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"up", "down"},
		Run:       run,
	}
}

func (m *Migration) main(args []string) {}
