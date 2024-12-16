package postgres

import (
	"github.com/mohammadne/takhir/pkg/databases/postgres"
	"go.uber.org/zap"
)

type User interface {
}

func NewUser(lg *zap.Logger, postgres *postgres.Postgres) User {
	return &user{logger: lg, postgres: postgres}
}

type user struct {
	logger   *zap.Logger
	postgres *postgres.Postgres
}