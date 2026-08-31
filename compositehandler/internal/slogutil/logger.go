package slogutil

import (
	"log/slog"
	"os"
)

func NewLogger(file os.File) slog.Logger {
	var c CompositeHandler
	jsonHandler := slog.NewJSONHandler(&file, nil)
	textHandler := slog.NewTextHandler(os.Stdout, nil)
	c.Add(jsonHandler)
	c.Add(textHandler)
	return *slog.New(&c)
}
