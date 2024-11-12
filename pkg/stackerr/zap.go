package stackerr

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapFields(err error) []zapcore.Field {
	return []zapcore.Field{
		zap.Stringers("stacktrace", StackTrace(err)),
		zap.Error(err),
	}
}
