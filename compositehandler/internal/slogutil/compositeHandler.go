package slogutil

import (
	"context"
	"log/slog"
)

type CompositeHandler struct {
	Handlers []slog.Handler
}

func (c *CompositeHandler) Enabled(ctx context.Context, l slog.Level) bool {
	for _, v := range c.Handlers {
		if !v.Enabled(ctx, l) {
			return false
		}
	}
	return true
}

func (c *CompositeHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, v := range c.Handlers {
		err := v.Handle(ctx, r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CompositeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	for i, v := range c.Handlers {
		c.Handlers[i] = v.WithAttrs(attrs)
	}
	return c
}

func (c *CompositeHandler) WithGroup(name string) slog.Handler {
	for i, v := range c.Handlers {
		c.Handlers[i] = v.WithGroup(name)
	}
	return c
}
