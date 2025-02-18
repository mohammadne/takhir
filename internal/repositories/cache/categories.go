package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/mohammadne/zanbil/internal/entities"
	"github.com/mohammadne/zanbil/pkg/databases/redis"
	"go.uber.org/zap"
)

type Categories interface {
	AllCategories(ctx context.Context) ([]entities.Category, error)
	SetAllCategories(ctx context.Context, categories []entities.Category) error
}

func NewCategories(logger *zap.Logger, redis *redis.Redis) Categories {
	return &categories{logger: logger, redis: redis}
}

type categories struct {
	logger *zap.Logger
	redis  *redis.Redis
}

const categoriesKey = "categories:all" // Key that stores all categories

var (
	ErrCategoriesNotFound            = errors.New("err_categories_not_found")
	ErrRetrievingCategoriesFromRedis = errors.New("err_retrieving_categories_from_redis")
	ErrUnmarshallingCachedCategories = errors.New("err_unmarshalling_cached_categories")
)

func (c *categories) AllCategories(ctx context.Context) ([]entities.Category, error) {
	cachedCategories, err := c.redis.Get(ctx, categoriesKey).Result()
	if errors.Is(err, redis.Nil) || len(cachedCategories) == 0 {
		return nil, ErrCategoriesNotFound
	} else if err != nil {
		c.logger.Error(ErrRetrievingCategoriesFromRedis.Error(), zap.Error(err))
		return nil, ErrRetrievingCategoriesFromRedis
	}

	categories := make([]entities.Category, 0, len(cachedCategories))
	err = json.Unmarshal([]byte(cachedCategories), categories)
	if err != nil {
		c.logger.Error(ErrUnmarshallingCachedCategories.Error(), zap.Error(err))
		return nil, ErrUnmarshallingCachedCategories
	}

	return categories, nil
}

const (
	SetAllCategoriesTimeout = 1 * time.Second
)

var (
	ErrMarshallingCategories  = errors.New("err_marshalling_categories")
	ErrSetCategoriesIntoRedis = errors.New("err_set_categories_into_redis")
)

func (c *categories) SetAllCategories(ctx context.Context, categories []entities.Category) error {
	ctx, cf := context.WithTimeout(ctx, SetAllCategoriesTimeout)
	defer cf()

	marshaledCategories, err := json.Marshal(categories)
	if err != nil {
		c.logger.Error(ErrMarshallingCategories.Error(), zap.Any("categories", categories), zap.Error(err))
		return ErrMarshallingCategories
	}

	err = c.redis.Set(ctx, categoriesKey, marshaledCategories, time.Hour).Err()
	if err != nil {
		c.logger.Error(ErrSetCategoriesIntoRedis.Error(), zap.Error(err))
		return ErrSetCategoriesIntoRedis
	}

	return nil
}
