package main

import (
	"time"

	app "github.com/tulashvili/MyDailyControlQuestions"
)

func main() {
	time.Local = time.UTC
	app.RunApp()
}
