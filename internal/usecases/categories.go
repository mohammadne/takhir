package usecases

import (
	"context"
	"errors"

	"github.com/mohammadne/zanbil/internal/entities"
	"github.com/mohammadne/zanbil/internal/repositories/cache"
	"github.com/mohammadne/zanbil/internal/repositories/storage"
	"go.uber.org/zap"
)

type Categories interface {
	AllCategories(ctx context.Context) ([]entities.Category, error)
	ListCategoryProducts(ctx context.Context, id uint64, options *entities.ListOptions) ([]entities.Category, error)
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

var (
	ErrRetrievingCategories = errors.New("err_retrieving_categories")
)

func (c *categories) AllCategories(ctx context.Context) ([]entities.Category, error) {
	categories, failure := c.categoriesCache.AllCategories(ctx)
	if failure == nil {
		return categories, nil
	}

	storageCategories, failure := c.categoriesStorage.AllCategories(ctx)
	if failure != nil {
		c.logger.Error(ErrRetrievingCategories.Error(), zap.Error(failure))
		return nil, ErrRetrievingCategories
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

func (c *categories) ListCategoryProducts(ctx context.Context, id uint64,
	options *entities.ListOptions) ([]entities.Category, error) {
	return nil, nil
}
