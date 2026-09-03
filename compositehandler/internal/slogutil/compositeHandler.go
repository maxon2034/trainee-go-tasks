package slogutil

import (
	"context"
	"errors"
	"log/slog"
	"slices"
)

type CompositeHandler struct {
	handlers []slog.Handler
}

func NewCompositeHandler(handlers ...slog.Handler) CompositeHandler {
	return CompositeHandler{handlers: slices.Clone(handlers)}
}

func (c CompositeHandler) Enabled(ctx context.Context, l slog.Level) bool {
	for _, v := range c.handlers {
		if v.Enabled(ctx, l) {
			return true
		}
	}
	return false
}

func (c CompositeHandler) Handle(ctx context.Context, r slog.Record) error {
	var errs []error
	var rc slog.Record
	for _, v := range c.handlers {

		if !v.Enabled(ctx, rc.Level) {
			continue
		}

		rc = r.Clone()

		err := v.Handle(ctx, rc)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (c CompositeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var ca CompositeHandler
	if attrs == nil {
		return c
	}
	for _, v := range c.handlers {
		ca.handlers = append(ca.handlers, v.WithAttrs(slices.Clone(attrs)))
	}

	return &ca
}

func (c CompositeHandler) WithGroup(name string) slog.Handler {
	var cg CompositeHandler
	if name == "" {
		return c
	}
	for _, v := range c.handlers {
		cg.handlers = append(cg.handlers, v.WithGroup(name))
	}
	return &cg
}
