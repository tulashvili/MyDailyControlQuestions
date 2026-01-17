package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(databasePath string) (*sql.DB, error) {
	fmt.Printf("🔌 Соединение с базой %s установлено\n", databasePath) // change to log?
	return sql.Open("sqlite3", databasePath)

}

func CreateTable(conn *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS daily_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		category TEXT NOT NULL,
		question TEXT NOT NULL,
		answer INTEGER NOT NULL,
		answeredAt TEXT NOT NULL,

		UNIQUE (answeredAt)
	);
	`
	_, err := conn.Exec(query)
	fmt.Println("✅ Таблица daily_log готова к использованию") // change to log?
	return err
}
