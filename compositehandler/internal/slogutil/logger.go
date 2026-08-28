package slogutil

import (
	"log/slog"
	"os"
)

func NewLogger(file os.File) slog.Logger {
	var c CompositeHandler
	jsonHandler := slog.NewJSONHandler(&file, nil)
	textHandler := slog.NewTextHandler(os.Stdout, nil)
	c.Handlers = append(c.Handlers, jsonHandler)
	c.Handlers = append(c.Handlers, textHandler)
	return *slog.New(&c)
}
