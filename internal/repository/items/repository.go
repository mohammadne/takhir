package items

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mohammadne/takhir/internal/entities/helpers"
)

type Items interface {
	ListItems(ctx context.Context, pagination *helpers.Pagination) ([]Item, error)
}

func New(db *sqlx.DB) Items {
	return &items{db: db}
}

type items struct {
	db *sqlx.DB
}

func (i *items) ListItems(ctx context.Context, pagination *helpers.Pagination) ([]Item, error) {
	return nil, nil
}
