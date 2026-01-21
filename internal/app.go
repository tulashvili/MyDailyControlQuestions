package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tulashvili/MyDailyControlQuestions/internal/config"
	repository "github.com/tulashvili/MyDailyControlQuestions/internal/repo"
	"github.com/tulashvili/MyDailyControlQuestions/internal/service"
	"github.com/tulashvili/MyDailyControlQuestions/internal/ui"

	"github.com/tulashvili/MyDailyControlQuestions/pkg/db/sqlitedb"
)

type App struct {
	DB     *sql.DB
	Config config.Config
}

func NewApp(conf config.Config) {
	app := &App{
		Config: conf,
	}

	// Logic???
	questions := service.GetQuestions()
	answers := ui.AskQuestion(questions)

	// DB
	var err error
	app.DB, err = sqlitedb.InitDB(conf.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlitedb.CreateTable(app.DB); err != nil {
		log.Fatal(err)
	}

	for _, answer := range answers {
		if err := repository.InsertRow(app.DB, answer); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("✅ Данные успешно добавлены в таблицу daily_log") // change to log?

	// UI
	if err := ui.ShowDataOverPeriod(app.DB); err != nil {
		log.Panic(err)
	}
}
