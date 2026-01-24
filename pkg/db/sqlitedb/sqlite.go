package sqlitedb

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tulashvili/MyDailyControlQuestions/pkg/db"
)

// move this to migrate.go

func InitDB(path db.DataSource) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", path.SqlitePath)
	if err == nil {
		slog.Info(
			"Соединение с базой установлено",
			"sqlite_path", path.SqlitePath,
		)
	}
	return conn, err

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
	if err == nil {
		slog.Info("Таблица daily_log готова к использованию")
	}
	return err
}
