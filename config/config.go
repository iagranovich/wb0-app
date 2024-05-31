package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		slog.Error("config: cannot load .env file", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
