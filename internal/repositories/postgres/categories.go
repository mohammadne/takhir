package postgres

import (
	"context"
	"database/sql"
	"time"

	"go.uber.org/zap"

	"github.com/mohammadne/takhir/pkg/databases/postgres"
	"github.com/mohammadne/takhir/pkg/metrics"
	"github.com/mohammadne/takhir/pkg/stackerr"
)

// counterLabels := []string{"function", "status", "info"}
// r.counterVector = t.Metric.NewCounter("ports-http-repository", counterLabels)

// timeLabels := []string{"function", "status"}
// r.timeVector = t.Metric.NewHistogram("ports-http-repository", timeLabels)

// type CounterVector struct {
// 	vector *prometheus.CounterVec
// }

// func NewCounter(namespace string, subsystem string, name string, labels []string) *CounterVector {
// 	vector := prometheus.NewCounterVec(prometheus.CounterOpts{
// 		Help:      fmt.Sprintf("counter vector for %s", name),
// 		Namespace: namespace,
// 		Subsystem: subsystem,
// 		Name:      name,
// 	}, labels)

// 	prometheus.MustRegister(vector)
// 	return &CounterVector{vector}
// }

// func (vector *CounterVector) Increment(labels ...string) {
// 	vector.vector.WithLabelValues(labels...).Inc()
// }

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

func (c *categories) CreateCategory(ctx context.Context, categoryDAO *CategoryDAO) (categoryID int, err error) {
	defer func(start time.Time) {
		c.database.Vectors.Counter.WithLabelValues("categories", "create_category", metrics.StatusSuccess).Inc()
		// if err != nil {
		// 	b.metric.IncrementError("define", "banners", err.Error())
		// }
		// b.metric.ObserveResponseTime(time.Since(start), "define", "banners")
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
		return 0, stackerr.Wrap(err, "error creating category in database")
	}

	return categoryID, nil
}
