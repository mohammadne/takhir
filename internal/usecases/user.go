package usecases

import (
	"context"
	"fmt"

	"github.com/mohammadne/takhir/internal/entities"
	"github.com/mohammadne/takhir/internal/repositories/storage"
	"go.uber.org/zap"
)

type Carts interface {
	RegisterUserByPhone(ctx context.Context, phone entities.Phone) error
}

func NewUser(logger *zap.Logger, userStorage storage.User) Carts {
	return &user{userStorage: userStorage}
}

type user struct {
	userStorage storage.User
}

func (u *user) RegisterUserByPhone(ctx context.Context, phone entities.Phone) error {
	phone = phone.Uniform()
	if !phone.Validate() {
		return fmt.Errorf("invalid phone number has been given: %s", phone)
	}

	return nil
}
