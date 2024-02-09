package log

import (
	"context"

	"go.uber.org/zap"
)

var (
	KeyTrace    = "trace"
	KeySubject  = "subject"
	KeyObject   = "object"
	KeyInstance = "instance"
	KeyTask     = "task"
	KeyCost     = "cost"
	KeyResult   = "result"
	KeyStatus   = "status"
	KeyIP       = "ip"
	KeyInfo     = "info"
	KeyWatcher  = "watcher"
)

// C Parse context.
func C(ctx context.Context) *zapLogger {
	return std.C(ctx)
}

func (l *zapLogger) C(ctx context.Context) *zapLogger {
	lc := l.clone()

	if data := ctx.Value(KeyTrace); data != nil {
		lc.z = lc.z.With(zap.Any(KeyTrace, data))
	}

	if data := ctx.Value(KeySubject); data != nil {
		lc.z = lc.z.With(zap.Any(KeySubject, data))
	}

	if data := ctx.Value(KeyObject); data != nil {
		lc.z = lc.z.With(zap.Any(KeyObject, data))
	}

	if data := ctx.Value(KeyInstance); data != nil {
		lc.z = lc.z.With(zap.Any(KeyInstance, data))
	}

	if data := ctx.Value(KeyTask); data != nil {
		lc.z = lc.z.With(zap.Any(KeyTask, data))
	}

	if data := ctx.Value(KeyCost); data != nil {
		lc.z = lc.z.With(zap.Any(KeyCost, data))
	}

	if data := ctx.Value(KeyResult); data != nil {
		lc.z = lc.z.With(zap.Any(KeyResult, data))
	}

	if data := ctx.Value(KeyStatus); data != nil {
		lc.z = lc.z.With(zap.Any(KeyStatus, data))
	}

	if data := ctx.Value(KeyIP); data != nil {
		lc.z = lc.z.With(zap.Any(KeyIP, data))
	}

	if data := ctx.Value(KeyInfo); data != nil {
		lc.z = lc.z.With(zap.Any(KeyInfo, data))
	}

	if data := ctx.Value(KeyWatcher); data != nil {
		lc.z = lc.z.With(zap.Any(KeyWatcher, data))
	}

	return lc
}

func (l *zapLogger) clone() *zapLogger {
	lc := *l

	return &lc
}
