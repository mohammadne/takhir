package usecases

import (
	"context"
	"fmt"

	"github.com/mohammadne/takhir/internal/entities"
	"github.com/mohammadne/takhir/internal/repositories/postgres"
)

type user struct {
	postgre postgres.User
}

func (u *user) RegisterUserByPhone(ctx context.Context, phone entities.Phone) error {
	phone = phone.Uniform()
	if !phone.Validate() {
		return fmt.Errorf("invalid phone number has been given: %s", phone)
	}

	return nil
}
