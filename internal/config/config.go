package config

import (
	"github.com/mohammadne/takhir/internal/repository"
	"github.com/mohammadne/takhir/pkg/logger"
)

type Config struct {
	Repository *repository.Config `koanf:"repository"`
	Logger     *logger.Config     `koanf:"logger"`
}
