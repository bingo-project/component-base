package log

import (
	"context"

	"go.uber.org/zap"
)

var (
	XRequestIDKey = "X-Request-ID"
	XUsernameKey  = "X-Username"
)

// C Parse context.
func C(ctx context.Context) *zapLogger {
	return std.C(ctx)
}

func (l *zapLogger) C(ctx context.Context) *zapLogger {
	lc := l.clone()

	if requestID := ctx.Value(XRequestIDKey); requestID != nil {
		lc.z = lc.z.With(zap.Any(XRequestIDKey, requestID))
	}

	if userID := ctx.Value(XUsernameKey); userID != nil {
		lc.z = lc.z.With(zap.Any(XUsernameKey, userID))
	}

	return lc
}

func (l *zapLogger) clone() *zapLogger {
	lc := *l

	return &lc
}
