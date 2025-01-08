package logger

type LoggerType string

const (
	StandardLogger LoggerType = "std"
	GraylogLogger  LoggerType = "graylog"
	SentryLogger   LoggerType = "sentry"
)

type Config struct {
	Development bool           `koanf:"development"`
	Loggers     []LoggerType   `koanf:"loggers"`
	Graylog     *GraylogConfig `koanf:"graylog"`
	Sentry      *SentryConfig  `koanf:"sentry"`
}

type GraylogConfig struct {
	URI      string `koanf:"uri"`
	Facility string `koanf:"facility"`
}

type SentryConfig struct {
	URI  string            `koanf:"uri"`
	Tags map[string]string `koanf:"tags"`
}
