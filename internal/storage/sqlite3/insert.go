package sqlite3

import (
	"database/sql"

	"github.com/tulashvili/MyDailyControlQuestions/internal/model"
)

func InsertRows(conn *sql.DB, data model.Entry) error {
	query := `
	INSERT INTO daily_log (category, question, scale, completedAt)
	VALUES ($1, $2, $3, $4)
	;
	`
	_, err := conn.Exec(
		query,
		data.QuestionCategory,
		data.QuestionText,
		data.Answer,
		data.CompletedAt,
	)
	return err
}
