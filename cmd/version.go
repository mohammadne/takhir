package cmd

import (
	"runtime"

	"go.uber.org/zap"
)

// Default build-time variable.
// These values are overridden via ldflags
var (
	Version    = "unknown-version"
	GitCommit  = "unknown-commit"
	BuildTime  = "unknown-buildtime"
	APIVersion = "v0.1.0"
)

func BuildInfo() []zap.Field {

	return []zap.Field{
		zap.String("Version", Version),
		zap.String("API Version", APIVersion),
		zap.String("Go Version", runtime.Version()),
		zap.String("Git Commit", GitCommit),
		zap.String("Built At", BuildTime),
		zap.String("OS", runtime.GOOS),
		zap.String("Arch", runtime.GOARCH),
	}
}
