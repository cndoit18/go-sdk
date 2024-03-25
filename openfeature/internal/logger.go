package internal

import (
	"context"
	"log/slog"
)

var _ slog.Handler = Logger{}

type Logger struct{}

func (l Logger) Enabled(context.Context, slog.Level) bool {
	return true
}

func (l Logger) Handle(context.Context, slog.Record) error {
	return nil
}

func (l Logger) WithAttrs(attrs []slog.Attr) slog.Handler {
	return l
}

func (l Logger) WithGroup(name string) slog.Handler {
	return l
}
