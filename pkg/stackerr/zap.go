package stackerr

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapFields(err error) []zapcore.Field {
	s, ok := err.(Stakerr)
	if ok {
		return []zapcore.Field{
			zap.Stringers("stacktrace", s.StackTrace()),
			zap.Error(err),
		}
	}
	return []zapcore.Field{
		zap.Error(err),
	}
}
