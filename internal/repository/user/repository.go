package user

import (
	"context"

	"github.com/mohammadne/takhir/pkg/postgres"
	"go.uber.org/zap"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error

	FindUserById(ctx context.Context, id uint64) (*User, error)

	FindUserByEmail(ctx context.Context, email string) (*User, error)

	FindUserByEmailAndPassword(ctx context.Context, email, password string) (*User, error)

	// UpdateUser will only updates the first_name and last_name or password
	UpdateUser(ctx context.Context, user *User) error

	DeleteUser(ctx context.Context, user *User) error
}

func New(lg *zap.Logger, postgres *postgres.Postgres) *repository {
	return &repository{logger: lg, postgres: postgres}
}

type repository struct {
	logger   *zap.Logger
	postgres *postgres.Postgres
}
