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
	var rc slog.Record
	for _, v := range c.handlers {
		rc = r.Clone()
		if !v.Enabled(ctx, rc.Level) {
			continue
		}
		err := v.Handle(ctx, rc)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (c *CompositeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var ca CompositeHandler
	for _, v := range c.handlers {
		ca.handlers = append(ca.handlers, v.WithAttrs(attrs))
	}

	return &ca
}

func (c *CompositeHandler) WithGroup(name string) slog.Handler {
	var cg CompositeHandler
	for _, v := range c.handlers {
		cg.handlers = append(cg.handlers, v.WithGroup(name))
	}
	return &cg
}

func (c *CompositeHandler) Add(h slog.Handler) *CompositeHandler {
	c.handlers = append(c.handlers, h)
	return c
}
