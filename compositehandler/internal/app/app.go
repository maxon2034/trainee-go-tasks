package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/maxon2034/trainee-go-tasks/compositehandler/internal/slogutil"
)

func Run() {
	ctx := context.Background()

	file, err := os.OpenFile("log.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error in creating file:", err)
		return
	}
	defer file.Close()

	logger := slogutil.NewLogger(*file)
	logger.Log(ctx, slog.LevelDebug, "aaa", "key1", "val1")
	logger.Log(ctx, slog.LevelInfo, "bbb", "key2", "val2")
	logger.Log(ctx, slog.LevelWarn, "ccc", "key3", "val3")
	logger.Log(ctx, slog.LevelError, "ddd", "key4", "val4")
	logger.Log(ctx, slog.LevelInfo, "eee", "key5", "val5")
}
