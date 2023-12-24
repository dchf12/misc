package main

import (
	"log/slog"
	"os"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError})))
}

func main() {
	slog.Info("hello world")
	slog.Warn("warn", slog.String("key", "value"), slog.Int("int", 1))
	slog.Error("error", "name", "value")
}
