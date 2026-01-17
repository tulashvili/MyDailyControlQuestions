package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tulashvili/MyDailyControlQuestions/internal/config"
	repository "github.com/tulashvili/MyDailyControlQuestions/internal/repo"
	"github.com/tulashvili/MyDailyControlQuestions/internal/service"
	"github.com/tulashvili/MyDailyControlQuestions/internal/ui"

	"github.com/tulashvili/MyDailyControlQuestions/pkg/db"
)

const (
	sqliteStoragePath = "data/sqlite3.db"
)

type App struct {
	DB     *sql.DB
	Config config.Config
}

func RunApp() {

	questions := service.GetQuestions()
	answers := ui.AskQuestion(questions)

	// DB
	conn, err := db.InitDB(sqliteStoragePath)
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

	if err := ui.ShowDataOverPeriod(conn); err != nil {
		log.Panic(err)
	}
}
