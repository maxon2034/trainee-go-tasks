package app

import (
	"context"
	"log/slog"

	"github.com/maxon2034/trainee-go-tasks/compositehandler/internal/slogutil"
)

func Run(ctx context.Context, handlers ...slog.Handler) {
	logger := slogutil.NewLogger(handlers...)
	logger.Log(ctx, slog.LevelDebug, "aaa", "key1", "val1")
	logger.Log(ctx, slog.LevelInfo, "bbb", "key2", "val2")
	logger.Log(ctx, slog.LevelWarn, "ccc", "key3", "val3")
	logger.Log(ctx, slog.LevelError, "ddd", "key4", "val4")
	logger.Log(ctx, slog.LevelInfo, "eee", "key5", "val5")
}
