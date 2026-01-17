package main

import (
	"time"

	app "github.com/tulashvili/MyDailyControlQuestions/internal"
)

func main() {
	time.Local = time.UTC
	app.RunApp()
}
