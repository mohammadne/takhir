package config

import (
	"github.com/mohammadne/takhir/pkg/databases/postgres"
	"github.com/mohammadne/takhir/pkg/logger"
)

type Config struct {
	Logger   *logger.Config   `koanf:"logger"`
	Postgres *postgres.Config `koanf:"postgres"`
}
