package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const Nil = redis.Nil

type Redis struct {
	*redis.Client
}

func Open(config *Config) (*Redis, error) {
	r := Redis{}

	r.Client = redis.NewClient(&redis.Options{
		Addr:         config.Address,
		DB:           config.DB,
		Username:     config.Username,
		Password:     config.Password,
		PoolSize:     config.PoolSize,
		DialTimeout:  config.Timeout,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
	})

	pingCtx, cf := context.WithTimeout(context.Background(), config.Timeout)
	defer cf()

	err := r.Client.Ping(pingCtx).Err()
	if err != nil {
		return nil, fmt.Errorf("error pinging the redis client: %v", err)
	}

	return &r, nil
}
