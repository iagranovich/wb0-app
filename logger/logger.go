package logger

import (
	"log/slog"
	"os"
)

func Setup() *slog.Logger {
	logg := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logg)
	return logg
}
