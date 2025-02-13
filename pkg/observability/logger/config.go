package logger

type LoggerType string

const (
	StandardLogger LoggerType = "std"
	GraylogLogger  LoggerType = "graylog"
	SentryLogger   LoggerType = "sentry"
)

type Config struct {
	Development bool         `required:"true"`
	Loggers     []LoggerType `required:"true"`
	Graylog     *GraylogConfig
	Sentry      *SentryConfig
}

type GraylogConfig struct {
	URI      string `required:"true"`
	Facility string `required:"true"`
}

type SentryConfig struct {
	URI  string `required:"true"`
	Tags map[string]string
}
