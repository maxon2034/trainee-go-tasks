package slogutil

import (
	"context"
	"log/slog"
)

type CompositeHandler struct {
	handlers []slog.Handler
}

func (c *CompositeHandler) Enabled(ctx context.Context, l slog.Level) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	for _, v := range c.handlers {
		if v.Enabled(ctx, l) {
			return true
		}
	}
	return false
}

func (c *CompositeHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, v := range c.handlers {
		err := v.Handle(ctx, r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CompositeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	for i, v := range c.handlers {
		c.handlers[i] = v.WithAttrs(attrs)
	}
	return c
}

func (c *CompositeHandler) WithGroup(name string) slog.Handler {
	for i, v := range c.handlers {
		c.handlers[i] = v.WithGroup(name)
	}
	return c
}
