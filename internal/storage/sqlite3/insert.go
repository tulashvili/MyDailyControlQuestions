package sqlite3

import (
	"database/sql"

	"github.com/tulashvili/MyDailyControlQuestions/internal/models"
)

func InsertRow(conn *sql.DB, data models.UserAnswer) error {
	query := `
	INSERT INTO daily_log (category, question, answer, answeredAt)
	VALUES (?, ?, ?, ?)
	`
	_, err := conn.Exec(
		query,
		data.QuestionCategory,
		data.QuestionText,
		data.Answer,
		data.AnsweredAt,
	)
	return err
}
