package config

import (
	"github.com/mohammadne/takhir/pkg/logger"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
