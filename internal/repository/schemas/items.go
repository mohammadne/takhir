package schemas

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mohammadne/takhir/internal/entities/helpers"
)

type Items interface {
	ListItems(ctx context.Context, pagination *helpers.Pagination) ([]Item, error)
}

type Item struct {
}

type items struct {
	db *sqlx.DB
}

func NewItems(db *sqlx.DB) Items {
	return &items{db: db}
}

func (i *items) ListItems(ctx context.Context, pagination *helpers.Pagination) ([]Item, error) {
	return nil, nil
}

func (i *items) getListItemsQueryAndArgs(pagination *helpers.Pagination) (string, []any) {
	return "", nil
}
