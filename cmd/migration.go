package cmd

import (
	"context"

	"github.com/mohammadne/takhir/internal/config"
	"github.com/mohammadne/takhir/pkg/databases/postgres"
	"github.com/mohammadne/takhir/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Migration struct {
	config    *config.Config
	logger    *zap.Logger
	direction string
	postgres  *postgres.Postgres
}

const (
	MigrateDirectionUp   string = "up"
	MigrateDirectionDown string = "down"
)

func (m Migration) Command(ctx context.Context) *cobra.Command {
	run := func(_ *cobra.Command, args []string) {
		m.initialize(args)
		m.run()
	}

	return &cobra.Command{
		Use:       "migration",
		Short:     "run migrations against database",
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{MigrateDirectionUp, MigrateDirectionDown},
		Run:       run,
	}
}

func (m *Migration) initialize(args []string) {
	m.config = config.Load(true)
	m.logger = logger.NewZap(m.config.Logger)

	if len(args) != 1 {
		m.logger.Fatal("invalid arguments have been given", zap.Any("args", args))
	} else {
		m.direction = args[0]
	}

	postgres, err := postgres.Open(m.config.Postgres, "file://hacks/migrations")
	if err != nil {
		m.logger.Fatal("error connecting to postgresql database", zap.Error(err))
	}
	m.postgres = postgres

	m.logger.Info("migration has been fully initialized")
}

func (m *Migration) run() {
	var callsMigrator func(context.Context) error
	if m.direction == MigrateDirectionUp {
		callsMigrator = m.postgres.MigrateUp
	} else {
		callsMigrator = m.postgres.MigrateDown
	}

	if err := callsMigrator(context.Background()); err != nil {
		m.logger.Fatal("Error while migrating the database",
			zap.String("direction", m.direction),
			zap.Error(err))
	}

	m.logger.Info("database migration has been successfully done",
		zap.String("direction", m.direction))
}
