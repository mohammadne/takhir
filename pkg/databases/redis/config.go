package redis

import (
	"time"
)

type Config struct {
	Address  string        `required:"true"`
	Username string        `required:"true"`
	Password string        `required:"true"`
	DB       int           `required:"true"`
	Timeout  time.Duration `default:"5s" required:"false"`
	PoolSize int           `default:"10" required:"false"`
}
