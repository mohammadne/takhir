package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mohammadne/takhir/internal/repository/items"
	"go.uber.org/zap"
)

type Repository interface {
	MigrateUp(context.Context) error
	MigrateDown(context.Context) error

	items.Items
}

type repository struct {
	logger             *zap.Logger
	database           *sqlx.DB
	migrationDirectory string

	// repositories
	items.Items
}

const (
	driver      = "postgres"
	pingTimeout = time.Second * 20
)

func Connect(cfg *Config, lg *zap.Logger) (Repository, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database,
	)

	database, err := sqlx.Open(driver, connString)
	if err != nil {
		errString := "Error while opening connection to postgresql"
		lg.Error(errString, zap.Error(err))
		return nil, fmt.Errorf(errString)
	}

	database.SetMaxIdleConns(0)

	ctx, cf := context.WithTimeout(context.Background(), pingTimeout)
	defer cf()
	if err = database.PingContext(ctx); err != nil {
		errString := "Error while pinging database"
		lg.Error(errString, zap.Error(err))
		return nil, fmt.Errorf(errString)
	}

	r := &repository{logger: lg, database: database}
	r.migrationDirectory = "file://hacks/migrations"

	// initialize repositories
	r.Items = items.New(database)

	return r, nil
}

func (r *repository) MigrateUp(ctx context.Context) error {
	migrator := func(m *migrate.Migrate) error { return m.Up() }
	return r.migrate(r.migrationDirectory, migrator)
}

func (r *repository) MigrateDown(ctx context.Context) error {
	migrator := func(m *migrate.Migrate) error { return m.Down() }
	return r.migrate(r.migrationDirectory, migrator)
}

func (r *repository) migrate(source string, migrator func(*migrate.Migrate) error) error {
	instance, err := postgres.WithInstance(r.database.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Error creating migrate instance\n%v", err)
	}

	migration, err := migrate.NewWithDatabaseInstance(source, driver, instance)
	if err != nil {
		return fmt.Errorf("Error loading migration files\n%v", err)
	}

	if err := migrator(migration); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Error doing migrations\n%v", err)
	}

	return nil
}
