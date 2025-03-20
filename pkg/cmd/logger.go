package cmd

import (
	"log/slog"
)

func Logger() *slog.Logger {
	// Create a new slog handler with the default PTerm logger
	handler := NewSlogHandler()

	// Create a new slog logger with the handler
	return slog.New(handler)
}
