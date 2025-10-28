package main

import (
	"ha-rpi-monitoring/v0.1/lib/env"
	"log/slog"
	"os"
	"strings"
)

func getLogLevelFromEnv() slog.Level {
	levelStr := env.GetEnv("LOG_LEVEL", "info")
	switch strings.ToLower(levelStr) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo // valeur par défaut si non définie/incorrecte
	}
}

func initLogger() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: getLogLevelFromEnv(),
	})
	Logger := slog.New(handler)
	slog.SetDefault(Logger)
}

func init() {
	initLogger()
}
