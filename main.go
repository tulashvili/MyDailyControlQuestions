package main

import (
	"log"

	"github.com/tulashvili/MyDailyControlQuestions/internal/service"
	"github.com/tulashvili/MyDailyControlQuestions/internal/storage/sqlite3"
	"github.com/tulashvili/MyDailyControlQuestions/internal/ui"
)

const (
	sqliteStoragePath = "data/sqlite3.db"
)

func main() {
	questions := service.GetQuestions()
	answers := ui.AskQuestion(questions)
	// ui.PrintResult(answers)

	conn, err := sqlite3.InitDB(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to sqlite3:", err)
	}

	if err := sqlite3.CreateTable(conn); err != nil {
		log.Fatal(err)
	}

	for _, answer := range answers {
		if err := sqlite3.InsertRow(conn, answer); err != nil {
			log.Fatal(err)
		}
	}
	// if err := sqlite3.InsertRows(conn, answers); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("✅ Данные успешно добавлены в таблицу daily_log")

}
