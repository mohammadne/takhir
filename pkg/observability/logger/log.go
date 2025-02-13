package logger

import (
	"fmt"
	"net/url"
	"os"

	"github.com/TheZeroSlave/zapsentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

func New(cfg *Config) (*zap.Logger, error) {
	var cores []zapcore.Core

	level := zap.InfoLevel
	if cfg.Development {
		level = zap.DebugLevel
	}

	for _, logger := range cfg.Loggers {
		switch logger {
		case StandardLogger:
			core := standardCore(cfg, level)
			cores = append(cores, core)
		case GraylogLogger:
			core, err := graylogCore(cfg, level)
			if err != nil {
				return nil, err
			}
			cores = append(cores, core)
		case SentryLogger:
			core, err := sentryCore(cfg)
			if err != nil {
				return nil, err
			}
			cores = append(cores, core)
		}
	}

	return zap.New(zapcore.NewTee(cores...)), nil
}

func standardCore(cfg *Config, level zapcore.Level) zapcore.Core {
	writerSyncer := zapcore.Lock(os.Stdout)

	var encoderConfig zapcore.EncoderConfig
	if cfg.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Development {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	levelEnablerFunc := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return zapcore.Level(level) <= lvl
	})

	return zapcore.NewCore(encoder, writerSyncer, levelEnablerFunc)
}

func graylogCore(cfg *Config, level zapcore.Level) (zapcore.Core, error) {
	graylogURI, err := url.Parse(cfg.Graylog.URI)
	if err != nil {
		return nil, err
	}

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	var writerSyncer zapcore.WriteSyncer
	switch graylogURI.Scheme {
	case "udp":
		udpWriter, err := gelf.NewUDPWriter(graylogURI.Host)
		if err != nil {
			return nil, err
		}
		udpWriter.Facility = cfg.Graylog.Facility
		writerSyncer = zapcore.AddSync(udpWriter)
	case "tcp":
		tcpWriter, err := gelf.NewTCPWriter(graylogURI.Host)
		if err != nil {
			return nil, err
		}
		tcpWriter.Facility = cfg.Graylog.Facility
		writerSyncer = zapcore.AddSync(tcpWriter)
	}

	levelEnablerFunc := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return zapcore.Level(level) <= lvl
	})

	core := zapcore.NewCore(encoder, writerSyncer, levelEnablerFunc)
	err = core.Sync()
	if err != nil {
		return nil, fmt.Errorf("got error (%v) on creating graylog core", err)
	}

	return core, nil
}

func sentryCore(cfg *Config) (zapcore.Core, error) {
	zapConf := zapsentry.Configuration{
		Tags:  cfg.Sentry.Tags,
		Level: zapcore.ErrorLevel,
	}

	core, err := zapsentry.NewCore(zapConf, zapsentry.NewSentryClientFromDSN(cfg.Sentry.URI))
	if err != nil {
		return nil, fmt.Errorf("got error (%v) on creating sentry core", err)
	}

	return core, nil
}
