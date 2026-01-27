package main

import (
	"log/slog"
	"os"

	app "github.com/tulashvili/MyDailyControlQuestions/internal"
	"github.com/tulashvili/MyDailyControlQuestions/internal/config"
	"github.com/tulashvili/MyDailyControlQuestions/internal/ui/cli"
)

func main() {
	conf, err := config.NewConfig(true)
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	app, err := app.NewApp(*conf)
	if err != nil {
		slog.Error("failed to create app", slog.Any("error", err))
		os.Exit(1)
	}
	cli.Execute(app.DB)

}
