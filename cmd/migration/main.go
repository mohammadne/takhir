package main

import (
	"embed"
	"flag"
	"log/slog"
	"os"
	"strings"

	"github.com/mohammadne/takhir/cmd"
	"github.com/mohammadne/takhir/internal/config"
	"github.com/mohammadne/takhir/internal/core"
	"github.com/mohammadne/takhir/pkg/databases/postgres"
)

//go:embed schemas/*.sql
var files embed.FS

func main() {
	direction := flag.String("direction", "", "Either 'UP' or 'DOWN'")
	flag.Parse() // Parse the command-line flags

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo})))
	cmd.BuildInfo()

	cfg, err := config.LoadDefaults(true)
	if err != nil {
		slog.Error(`error loading configs`, `Err`, err)
		os.Exit(1)
	}

	db, err := postgres.Open(cfg.Postgres, core.Namespace, core.System)
	if err != nil {
		slog.Error(`error connecting to postgres database`, `Err`, err)
		os.Exit(1)
	}

	migrateDirection := postgres.MigrateDirection(strings.ToUpper(*direction))
	err = db.Migrate("schemas", &files, migrateDirection)
	if err != nil {
		slog.Error(`error migrating postgres database`, `Err`, err)
		os.Exit(1)
	}

	slog.Info(`database has been migrated`)
}
