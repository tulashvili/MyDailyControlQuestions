package main

import (
	"fmt"
	"log"

	"github.com/tulashvili/MyDailyControlQuestions/internal/service"
	"github.com/tulashvili/MyDailyControlQuestions/internal/storage"
	"github.com/tulashvili/MyDailyControlQuestions/internal/ui"
)

const (
	sqliteStoragePath = "data/sqlite3.db"
)

func main() {
	questions := service.GetQuestions()
	answers := ui.AskQuestion(questions)

	conn, err := storage.InitDB(sqliteStoragePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := storage.CreateTable(conn); err != nil {
		log.Fatal(err)
	}

	for _, answer := range answers {
		if err := storage.InsertRow(conn, answer); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("✅ Данные успешно добавлены в таблицу daily_log") // change to log?

	storage.SelectRows(conn, 2)
	if err != nil {
		log.Fatal("error select data from db", err)
	}

}
