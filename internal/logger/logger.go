package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
)

func NewLogger(isProd bool, logFile string) *slog.Logger {
	var writers []io.Writer
	writers = append(writers, os.Stdout)

	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}

		writers = append(writers, file)
	}

	multi := io.MultiWriter(writers...)

	if isProd {
		return slog.New(slog.NewJSONHandler(multi, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return slog.New(slog.NewTextHandler(multi, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}
