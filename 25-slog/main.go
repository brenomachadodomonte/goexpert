package main

import (
	"log/slog"
	"os"
)

func main() {

	slog.Info("hello", "count", 3)
	slog.Error("Testando", "idade", 30)

	var programLevel = new(slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)

	slog.Group("request",
		"method", "GET",
		"url", "/pessoa/10")
}
