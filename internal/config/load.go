package config

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
)

const (
	delimiter = "."
	tagName   = "koanf"

	upTemplate     = "================ Loaded Configuration ================"
	bottomTemplate = "======================================================"
)

func Load(print bool) *Config {
	k := koanf.New(delimiter)

	// load default configuration from defaults file
	if err := loadDefaults(k); err != nil {
		log.Fatalf("Error loading default values: \n%v", err)
	}

	// load config from configmap
	if err := loadConfigmap(k); err != nil {
		log.Fatalf("Error loading from configmap: \n%v", err)
	}

	config := Config{}
	var tag = koanf.UnmarshalConf{Tag: tagName}
	if err := k.UnmarshalWithConf("", &config, tag); err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

	if print {
		// pretty print loaded configuration using provided template
		log.Printf("%s\n%v\n%s\n", upTemplate, spew.Sdump(config), bottomTemplate)
	}

	return &config
}

//go:embed default.yml
var defaults []byte

// loadDefaults loads the configuration from the defaults file
func loadDefaults(k *koanf.Koanf) error {
	if err := k.Load(rawbytes.Provider(defaults), yaml.Parser()); err != nil {
		return fmt.Errorf("Error loading default values: %s", err)
	}

	return nil
}

const (
	envPrefix    = "GORILLAMQ"
	envSeperator = "__"
)

// loadEnv loads the configuration from environment variables
func loadEnv(k *koanf.Koanf) error {
	var prefix = envPrefix + envSeperator

	callback := func(source string) string {
		base := strings.ToLower(strings.TrimPrefix(source, prefix))
		return strings.ReplaceAll(base, envSeperator, delimiter)
	}

	// load environment variables
	if err := k.Load(env.Provider(prefix, delimiter, callback), nil); err != nil {
		return fmt.Errorf("error loading environment variables: %s", err)
	}

	return nil
}

func loadConfigmap(k *koanf.Koanf) error {
	if os.Getenv("RUNNING_INSIDE_POD") == "" {
		return nil
	}

	cm, err := os.ReadFile("/etc/lionmq/config.yaml")
	if err != nil {
		return fmt.Errorf("Error reading currnet namespace: %v", err)
	}

	if err := k.Load(rawbytes.Provider(cm), yaml.Parser()); err != nil {
		return fmt.Errorf("Error loading values: %s", err)
	}

	return nil
}
