package logger

import (
	"log/slog"
	"os"
)

func NewLogger(isProd bool) *slog.Logger {
	var handler slog.Handler

	if isProd {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	return slog.New(handler)
}
