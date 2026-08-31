package slogutil

import (
	"context"
	"errors"
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
	var errs []error
	for _, v := range c.handlers {
		if !v.Enabled(ctx, r.Level) {
			continue
		}
		err := v.Handle(ctx, r)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
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
