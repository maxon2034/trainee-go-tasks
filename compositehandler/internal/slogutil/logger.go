package slogutil

import (
	"io"
	"log/slog"
)

func NewLogger(w io.Writer) *slog.Logger {
	var c CompositeHandler
	jsonHandler := slog.NewJSONHandler(w, nil)
	textHandler := slog.NewTextHandler(w, nil)
	c.Add(jsonHandler)
	c.Add(textHandler)
	return slog.New(&c)
}
