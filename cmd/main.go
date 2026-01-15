package main

import (
	"fmt"
	"log"

	"github.com/tulashvili/MyDailyControlQuestions/internal/storage"
	"github.com/tulashvili/MyDailyControlQuestions/internal/ui"
)

const (
	sqliteStoragePath = "data/sqlite3.db"
)

func main() {
	// questions := service.GetQuestions()
	// answers := ui.AskQuestion(questions)

	conn, err := storage.InitDB(sqliteStoragePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := storage.CreateTable(conn); err != nil {
		log.Fatal(err)
	}

	// for _, answer := range answers {
	// 	if err := storage.InsertRow(conn, answer); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	fmt.Println("✅ Данные успешно добавлены в таблицу daily_log") // change to log?

	if err := ui.ShowDataOverPeriod(conn); err != nil {
		log.Panic(err)
	}
}
