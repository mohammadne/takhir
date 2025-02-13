package postgres

import (
	"context"
	"database/sql"
	"time"

	"go.uber.org/zap"

	"github.com/mohammadne/takhir/pkg/databases/postgres"
	"github.com/mohammadne/takhir/pkg/observability/metrics"
	// "github.com/mohammadne/takhir/pkg/stackerr"
)

type Categories interface {
}

type CategoryDAO struct {
	ID          int
	Name        string
	Description string
}

func NewCategories(lg *zap.Logger, postgres *postgres.Postgres) Categories {
	return &categories{
		logger:   lg,
		database: postgres,
	}
}

type categories struct {
	logger   *zap.Logger
	database *postgres.Postgres
}

func (c *categories) CreateOne(ctx context.Context, categoryDAO *CategoryDAO) (categoryID int, err error) {
	defer func(start time.Time) {
		if err != nil {
			c.database.Vectors.Counter.IncrementVector("categories", "create_category", metrics.StatusFailure)
			return
		}
		c.database.Vectors.Counter.IncrementVector("categories", "create_category", metrics.StatusSuccess)
		c.database.Vectors.Histogram.ObserveResponseTime(start, "categories", "create_category")
	}(time.Now())

	query := `
	INSERT INTO categories (name, description)
	VALUES (:name, :description)
	RETURNING id INTO :id`

	_, err = c.database.ExecContext(ctx, query,
		sql.Named("name", categoryDAO.Name),
		sql.Named("description", categoryDAO.Description),
		sql.Named("id", sql.Out{
			Dest: &categoryID,
		}),
	)
	if err != nil {
		// return 0, stackerr.Wrap(err, "error creating category in database")
	}

	return categoryID, nil
}
