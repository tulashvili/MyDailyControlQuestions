package main

import (
	"time"

	app "github.com/tulashvili/MyDailyControlQuestions/internal"
	"github.com/tulashvili/MyDailyControlQuestions/internal/ui/cli"
)

func main() {
	time.Local = time.UTC
	app.NewApp()

	cli.Execute()
}
