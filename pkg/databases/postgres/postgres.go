package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/mohammadne/takhir/pkg/observability/metrics"
)

type Postgres struct {
	*sqlx.DB
	Vectors *vectors
}

type vectors struct {
	Counter   metrics.Counter
	Histogram metrics.Histogram
}

const (
	driver      = "postgres"
	pingTimeout = time.Second * 20
)

func Open(cfg *Config, namespace, subsystem string) (*Postgres, error) {
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
	vectors.Counter, err = metrics.RegisterCounter(name, namespace, subsystem, counterLabels)
	if err != nil {
		return nil, fmt.Errorf("error while registering counter vector: %v", err)
	}

	histogramLabels := []string{"table, function"}
	vectors.Histogram, err = metrics.RegisterHistogram(name, namespace, subsystem, histogramLabels)
	if err != nil {
		return nil, fmt.Errorf("error while registering histogram vector: %v", err)
	}

	r := &Postgres{DB: database, Vectors: &vectors}

	return r, nil
}
