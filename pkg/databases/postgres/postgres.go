package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mohammadne/takhir/internal"
	"github.com/mohammadne/takhir/pkg/metrics"
)

type Postgres struct {
	*sqlx.DB
	migrations string
	Vectors    *vectors
}

type vectors struct {
	Counter   metrics.Counter
	Histogram metrics.Histogram
}

const (
	driver      = "postgres"
	pingTimeout = time.Second * 20
)

func Open(cfg *Config, migrations string) (*Postgres, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database,
	)

	database, err := sqlx.Open(driver, connString)
	if err != nil {
		return nil, fmt.Errorf("error while opening connection to postgresql: %v", err)
	}

	database.SetMaxIdleConns(0)

	ctx, cf := context.WithTimeout(context.Background(), pingTimeout)
	defer cf()
	if err = database.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error while pinging database: %v", err)
	}

	var vectors vectors
	name := "postgres"

	counterLabels := []string{"table, function", "status"}
	vectors.Counter, err = metrics.RegisterCounter(name, internal.Namespace, internal.Namespace, counterLabels)
	if err != nil {
		return nil, fmt.Errorf("error while registering counter vector: %v", err)
	}

	histogramLabels := []string{"table, function"}
	vectors.Histogram, err = metrics.RegisterHistogram(name, internal.Namespace, internal.Namespace, histogramLabels)
	if err != nil {
		return nil, fmt.Errorf("error while registering histogram vector: %v", err)
	}

	r := &Postgres{DB: database, migrations: migrations, Vectors: &vectors}

	return r, nil
}

func (r *Postgres) MigrateUp(ctx context.Context) error {
	migrator := func(m *migrate.Migrate) error { return m.Up() }
	return r.migrate(r.migrations, migrator)
}

func (r *Postgres) MigrateDown(ctx context.Context) error {
	migrator := func(m *migrate.Migrate) error { return m.Down() }
	return r.migrate(r.migrations, migrator)
}

func (r *Postgres) migrate(source string, migrator func(*migrate.Migrate) error) error {
	instance, err := postgres.WithInstance(r.DB.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creating migrate instance\n%v", err)
	}

	migration, err := migrate.NewWithDatabaseInstance(source, driver, instance)
	if err != nil {
		return fmt.Errorf("error loading migration files\n%v", err)
	}

	if err := migrator(migration); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error doing migrations\n%v", err)
	}

	return nil
}
