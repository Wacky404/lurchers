package util

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type MultiHandler struct {
	handlers []slog.Handler
}

func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m *MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, h := range m.handlers {
		_ = h.Handle(ctx, record) // Process each handler, ignoring errors
	}
	return nil
}

func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithAttrs(attrs)
	}
	return &MultiHandler{handlers: newHandlers}
}

func (m *MultiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithGroup(name)
	}
	return &MultiHandler{handlers: newHandlers}
}

type options struct {
	logName     string
	stdoutLevel slog.Level
	fileLevel   slog.Level
}

func SetupLogger(opts ...func(*options)) (*os.File, error) {
	o := options{
		logName:     "logs/Default.log",
		stdoutLevel: slog.LevelInfo,
		fileLevel:   slog.LevelDebug,
	}

	for _, opt := range opts {
		opt(&o)
	}

	logFile, err := os.OpenFile(o.logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil, fmt.Errorf("no file opened for logs: %v", err)
	}

	stdoutHandler := slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{Level: o.stdoutLevel})
	fileHandler := slog.NewJSONHandler(logFile,
		&slog.HandlerOptions{Level: o.fileLevel})

	logger := slog.New(NewMultiHandler(stdoutHandler, fileHandler))
	slog.SetDefault(logger)

	return logFile, err
}

func WithLogName(name string) func(*options) {
	return func(o *options) {
		o.logName = name
	}
}

func WithStdoutLevel(level slog.Level) func(*options) {
	return func(o *options) {
		o.stdoutLevel = level
	}
}

func WithFileLevel(level slog.Level) func(*options) {
	return func(o *options) {
		o.fileLevel = level
	}
}
