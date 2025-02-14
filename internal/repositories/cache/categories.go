package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/mohammadne/takhir/internal/entities"
	"github.com/mohammadne/takhir/pkg/databases/redis"
	"go.uber.org/zap"
)

type Categories interface {
	AllCategories(ctx context.Context) ([]entities.Category, entities.Failure)
	SetAllCategories(ctx context.Context, categories []entities.Category) entities.Failure
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
	FailureCategoriesNotFound            = entities.NewFailure("no_categories_have_been_found")
	FailureRetrievingCategoriesFromRedis = entities.NewFailure("failure_retrieving_categories_from_redis")
	FailureUnmarshallingCachedCategories = entities.NewFailure("failure_unmarshalling_cached_categories")
)

func (c *categories) AllCategories(ctx context.Context) ([]entities.Category, entities.Failure) {
	cachedCategories, err := c.redis.Get(ctx, categoriesKey).Result()
	if errors.Is(err, redis.Nil) || len(cachedCategories) == 0 {
		return nil, FailureCategoriesNotFound
	} else if err != nil {
		c.logger.Error(FailureRetrievingCategoriesFromRedis.Error(), zap.Error(err))
		return nil, FailureRetrievingCategoriesFromRedis
	}

	categories := make([]entities.Category, 0, len(cachedCategories))
	err = json.Unmarshal([]byte(cachedCategories), categories)
	if err != nil {
		c.logger.Error(FailureUnmarshallingCachedCategories.Error(), zap.Error(err))
		return nil, FailureUnmarshallingCachedCategories
	}

	return categories, nil
}

const (
	SetAllCategoriesTimeout = 1 * time.Second
)

var (
	FailureMarshallingCategories  = entities.NewFailure("failure_marshalling_categories")
	FailureSetCategoriesIntoRedis = entities.NewFailure("failure_set_categories_into_redis")
)

func (c *categories) SetAllCategories(ctx context.Context, categories []entities.Category) entities.Failure {
	ctx, cf := context.WithTimeout(ctx, SetAllCategoriesTimeout)
	defer cf()

	marshaledCategories, err := json.Marshal(categories)
	if err != nil {
		c.logger.Error(FailureMarshallingCategories.Error(), zap.Any("categories", categories), zap.Error(err))
		return FailureMarshallingCategories
	}

	err = c.redis.Set(ctx, categoriesKey, marshaledCategories, time.Hour).Err()
	if err != nil {
		c.logger.Error(FailureSetCategoriesIntoRedis.Error(), zap.Error(err))
		return FailureSetCategoriesIntoRedis
	}

	return nil
}
