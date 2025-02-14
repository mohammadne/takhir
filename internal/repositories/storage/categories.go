package storage

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/mohammadne/takhir/internal/entities"
	"github.com/mohammadne/takhir/pkg/databases/postgres"
	"github.com/mohammadne/takhir/pkg/observability/metrics"
)

type Categories interface {
	AllCategories(ctx context.Context) ([]Category, entities.Failure)
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
	FailureCategoriesNotFound       = entities.NewFailure("failure_categories_not_found")
	FailureRetrievingCategoriesRows = entities.NewFailure("failure_retrieving_categories_rows")
	FailureScanningCategoryRow      = entities.NewFailure("failure_retrieving_categories_rows")
)

func (c *categories) AllCategories(ctx context.Context) (result []Category, f entities.Failure) {
	defer func(start time.Time) {
		if f != nil {
			c.postgres.Vectors.Counter.IncrementVector("categories", "all_categories", metrics.StatusFailure)
			return
		}
		c.postgres.Vectors.Counter.IncrementVector("categories", "all_categories", metrics.StatusSuccess)
		c.postgres.Vectors.Histogram.ObserveResponseTime(start, "categories", "all_categories")
	}(time.Now())

	query := `
	SELECT id, name, description, created_at
	FROM categories`

	err := c.postgres.SelectContext(ctx, &result, query)
	if err != nil {
		c.logger.Error(FailureRetrievingCategoriesRows.Error(), zap.Error(err))
		return nil, FailureRetrievingCategoriesRows
	} else if len(result) == 0 {
		return nil, FailureCategoriesNotFound
	}

	return result, nil
}
