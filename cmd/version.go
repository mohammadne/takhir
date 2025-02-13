package cmd

import (
	"fmt"
	"runtime"
)

// Default build-time variable.
// These values are overridden via ldflags
var (
	Version    = "unknown-version"
	GitCommit  = "unknown-commit"
	BuildTime  = "unknown-buildtime"
	APIVersion = "v0.1.0"
)

func BuildInfo() map[string]string {
	return map[string]string{
		`Version`:     Version,
		`API Version`: APIVersion,
		`Go Version`:  runtime.Version(),
		`Git Commit`:  GitCommit,
		`Built At`:    BuildTime,
		`OS/ARCH`:     fmt.Sprintf("OS/Arch:\t %s/%s\n", runtime.GOOS, runtime.GOARCH),
	}
}
