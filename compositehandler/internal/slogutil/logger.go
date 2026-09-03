package slogutil

import "log/slog"

func NewLogger(handlers ...slog.Handler) slog.Logger {
	return *slog.New(NewCompositeHandler(handlers...))
}
