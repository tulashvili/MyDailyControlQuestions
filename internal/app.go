package app

import (
	"fmt"
	"log"

	"github.com/tulashvili/MyDailyControlQuestions/internal/config"
	repository "github.com/tulashvili/MyDailyControlQuestions/internal/repo"
	"github.com/tulashvili/MyDailyControlQuestions/internal/service"
	"github.com/tulashvili/MyDailyControlQuestions/internal/ui"

	"github.com/tulashvili/MyDailyControlQuestions/pkg/db"
)

type App struct {
	Config config.Config
}

func RunApp() {
	// app := App{
	// 	Config: cfg,
	// }
	
	// config.LoadConfig()

	// Logic???
	questions := service.GetQuestions()
	answers := ui.AskQuestion(questions)
	dbPath := db.LoadDatasource()

	// DB
	conn, err := db.InitDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.CreateTable(conn); err != nil {
		log.Fatal(err)
	}

	for _, answer := range answers {
		if err := repository.InsertRow(conn, answer); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("✅ Данные успешно добавлены в таблицу daily_log") // change to log?

	// UI
	if err := ui.ShowDataOverPeriod(conn); err != nil {
		log.Panic(err)
	}
}
