package logger

import (
	"log/slog"
	"os"
)

func SetupLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
}
