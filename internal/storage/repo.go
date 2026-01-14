package storage

import (
	"database/sql"

	"github.com/tulashvili/MyDailyControlQuestions/internal/models"
)

type UserAnswerRow struct {
	ID         int
	Category   string
	Question   string
	Answer     int
	AnsweredAt string
}

// ToDo: Create interface for db methods ??
// type .... interface { .... }

// Insert question answers
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

// Get data over some period
func SelectRows(conn *sql.DB) ([]UserAnswerRow, error) {
	query := `
	SELECT * FROM daily_log
	WHERE answeredAt >= datetime('now', '-7 days');
	`

	response, err := conn.Query(query)

	currentDay := make([]UserAnswerRow, 0) // rename variable?

	for response.Next() {
		var row UserAnswerRow

		err := response.Scan(
			&row.ID,
			&row.Category,
			&row.Question,
			&row.Answer,
			&row.AnsweredAt,
		)
		if err != nil {
			return nil, err
		}
		currentDay = append(currentDay, row)
	}
	return currentDay, err
}
