package config

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"

	"github.com/mohammadne/zanbil/internal/core"
	"github.com/mohammadne/zanbil/pkg/databases/postgres"
	"github.com/mohammadne/zanbil/pkg/databases/redis"
	"github.com/mohammadne/zanbil/pkg/observability/logger"
)

type Config struct {
	Logger   *logger.Config   `required:"true"`
	Postgres *postgres.Config `required:"true"`
	Redis    *redis.Config    `required:"true"`
}

func Load(print bool) (config Config, err error) {
	prefix := strings.ToUpper(core.System)

	if err = envconfig.Process(prefix, &config); err != nil {
		return Config{}, fmt.Errorf("error processing config via envconfig: %v", err)
	}

	if print {
		fmt.Println("================ Loaded Configuration ================")
		object, _ := json.MarshalIndent(config, "", "  ")
		fmt.Println(string(object))
		fmt.Println("======================================================")
	}

	return config, nil
}

const seperator = "_"

//go:embed defaults.env
var defaults string

func LoadDefaults(print bool) (config Config, err error) {
	lines := strings.Split(defaults, "\n")
	for _, line := range lines {
		splits := strings.Split(line, "=")
		if len(splits) < 2 {
			continue
		}

		key := strings.ReplaceAll(splits[0], seperator+seperator, seperator)
		err = os.Setenv(key, splits[1])
		if err != nil {
			return Config{}, fmt.Errorf("error set environment %s: %v", key, err)
		}
	}

	return Load(print)
}
