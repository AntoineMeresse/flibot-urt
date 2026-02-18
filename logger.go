package main

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func configureLogger() {
	levelStr := os.Getenv("LOG_LEVEL")

	var level slog.Level
	if err := level.UnmarshalText([]byte(levelStr)); err != nil {
		level = slog.LevelDebug
	}

	h := tint.NewHandler(os.Stdout, &tint.Options{Level: level})
	slog.SetDefault(slog.New(h))
}
