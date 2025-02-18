package storage

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/mohammadne/zanbil/pkg/databases/postgres"
	"github.com/mohammadne/zanbil/pkg/observability/metrics"
)

type Categories interface {
	AllCategories(ctx context.Context) ([]Category, error)
}

type Category struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

func NewCategories(lg *zap.Logger, postgres *postgres.Postgres) Categories {
	return &categories{
		logger:   lg,
		postgres: postgres,
	}
}

type categories struct {
	logger   *zap.Logger
	postgres *postgres.Postgres
}

var (
	ErrCategoriesNotFound       = errors.New("err_categories_not_found")
	ErrRetrievingCategoriesRows = errors.New("err_retrieving_categories_rows")
)

func (c *categories) AllCategories(ctx context.Context) (result []Category, err error) {
	defer func(start time.Time) {
		if err != nil {
			c.postgres.Vectors.Counter.IncrementVector("categories", "all_categories", metrics.StatusFailure)
			return
		}
		c.postgres.Vectors.Counter.IncrementVector("categories", "all_categories", metrics.StatusSuccess)
		c.postgres.Vectors.Histogram.ObserveResponseTime(start, "categories", "all_categories")
	}(time.Now())

	query := `
	SELECT id, name, description, created_at
	FROM categories`

	if err = c.postgres.SelectContext(ctx, &result, query); err != nil {
		c.logger.Error(ErrRetrievingCategoriesRows.Error(), zap.Error(err))
		return nil, ErrRetrievingCategoriesRows
	} else if len(result) == 0 {
		return nil, ErrCategoriesNotFound
	}

	return result, nil
}
