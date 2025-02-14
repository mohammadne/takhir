package usecases

import (
	"context"

	"github.com/mohammadne/takhir/internal/entities"
	"github.com/mohammadne/takhir/internal/repositories/cache"
	"github.com/mohammadne/takhir/internal/repositories/storage"
	"go.uber.org/zap"
)

type Categories interface {
	AllCategories(ctx context.Context) ([]entities.Category, entities.Failure)
}

func NewCategories(logger *zap.Logger,
	categoriesCache cache.Categories,
	categoriesStorage storage.Categories,
) Categories {
	return &categories{
		logger:            logger,
		categoriesCache:   categoriesCache,
		categoriesStorage: categoriesStorage,
	}
}

type categories struct {
	logger *zap.Logger

	// cache
	categoriesCache cache.Categories

	// storages
	categoriesStorage storage.Categories
}

func (c *categories) AllCategories(ctx context.Context) ([]entities.Category, entities.Failure) {
	categories, failure := c.categoriesCache.AllCategories(ctx)
	if failure == nil {
		return categories, nil
	}

	storageCategories, failure := c.categoriesStorage.AllCategories(ctx)
	if failure != nil {
		return nil, nil
	}

	categories = make([]entities.Category, 0, len(storageCategories))
	for _, storageCategory := range storageCategories {
		categories = append(categories, entities.Category{
			ID:          storageCategory.ID,
			Name:        storageCategory.Name,
			Description: storageCategory.Description,
		})
	}

	go c.categoriesCache.SetAllCategories(context.Background(), categories)

	return categories, nil
}
