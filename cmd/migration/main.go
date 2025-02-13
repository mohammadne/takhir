package main

import (
	"embed"
	"flag"
	"log"
	"strings"

	"github.com/mohammadne/takhir/internal/config"
	"github.com/mohammadne/takhir/internal/core"
	"github.com/mohammadne/takhir/pkg/databases/postgres"
)

//go:embed schemas/*.sql
var files embed.FS

func main() {
	direction := flag.String("direction", "", "Either 'UP' or 'DOWN'")
	environmentRaw := flag.String("environment", "", "The environment (default: local)")
	flag.Parse() // Parse the command-line flags

	var cfg config.Config
	var err error

	switch core.ToEnvironment(*environmentRaw) {
	case core.EnvironmentLocal:
		cfg, err = config.LoadDefaults(true)
	default:
		cfg, err = config.Load(true)
	}

	if err != nil {
		log.Fatalf("failed to load config: \n%v", err)
	}

	db, err := postgres.Open(cfg.Postgres, core.Namespace, core.System)
	if err != nil {
		log.Fatalf("error connecting to postgres database\n%v", err)
	}

	migrateDirection := postgres.MigrateDirection(strings.ToUpper(*direction))
	err = db.Migrate("schemas", &files, migrateDirection)
	if err != nil {
		log.Fatalf("error migrating postgres database\n%v", err)
	}

	log.Println("database has been migrated")
}
