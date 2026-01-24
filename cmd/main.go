package main

import (
	"log/slog"
	"os"

	app "github.com/tulashvili/MyDailyControlQuestions/internal"
	"github.com/tulashvili/MyDailyControlQuestions/internal/config"
)

func main() {
	conf, err := config.NewConfig(true)
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	err = app.NewApp(*conf)
	if err != nil {
		slog.Error("failed to create app", slog.Any("error", err))
		os.Exit(1)
	}

}
