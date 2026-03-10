package logger

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(isProd bool, logFile string) *slog.Logger {
	var writers []io.Writer
	writers = append(writers, os.Stdout)

	if logFile != "" {
		writers = append(writers, &lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    20,
			MaxBackups: 14,
			MaxAge:     14,
			Compress:   false,
		})
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
