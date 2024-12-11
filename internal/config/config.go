package config

import (
	"github.com/mohammadne/takhir/pkg/logger"
	"github.com/mohammadne/takhir/pkg/postgres"
)

type Config struct {
	Logger   *logger.Config   `koanf:"logger"`
	Postgres *postgres.Config `koanf:"postgres"`
}
