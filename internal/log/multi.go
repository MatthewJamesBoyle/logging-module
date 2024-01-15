package log

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type Logger interface {
	InfoContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
}

type MultiSourceLogger struct {
	logger []Logger
}

func NewMultiSourceLoggerLogger(opts *slog.HandlerOptions, writers ...io.Writer) *MultiSourceLogger {

	loggers := make([]Logger, 0)
	for _, v := range writers {
		loggers = append(loggers, slog.New(slog.NewJSONHandler(v, opts)))
	}
	loggers = append(loggers, slog.New(slog.NewJSONHandler(os.Stdout, opts)))

	return &MultiSourceLogger{
		logger: loggers,
	}
}

func (e MultiSourceLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	for _, v := range e.logger {
		v.InfoContext(ctx, msg, args...)
	}
}

func (e MultiSourceLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	for _, v := range e.logger {
		v.ErrorContext(ctx, msg, args...)
	}
}
