package sqlite3

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenConnection(s SqliteConf) (*sql.DB, error) {
	fmt.Printf("🔌 Соединение с базой %s установлено\n", s.PathToDB)
	return sql.Open("sqlite3", s.PathToDB)

}

func CreateTable(conn *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS daily_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		category TEXT NOT NULL,
		question TEXT NOT NULL,
		scale INTEGER NOT NULL,
		completedAt TEXT NOT NULL,

		UNIQUE (completedAt)
	);
	`
	_, err := conn.Exec(query)
	fmt.Println("✅ Таблица daily_log готова к использованию")
	return err
}
