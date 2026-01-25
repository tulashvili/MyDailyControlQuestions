package repo

import (
	"database/sql"
	"fmt"

	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
)

type UserAnswerRow struct {
	ID         int
	Category   string
	Question   string
	Answer     int
	AnsweredAt string
}

type SQLiteRepo struct {
    db *sql.DB
}

// ToDo: Create interface for db methods ??
// type .... interface { .... }

// Insert question answers
func InsertRow(conn SQLiteRepo, data domain.UserAnswer) error {
	query := `
	INSERT INTO daily_log (category, question, answer, answeredAt)
	VALUES (?, ?, ?, ?)
	`
	_, err := conn.db.Exec(
		query,
		data.QuestionCategory,
		data.QuestionText,
		data.Answer,
		data.AnsweredAt,
	)
	return err
}

// Get data over some period
func SelectRows(conn SQLiteRepo, period int) ([]UserAnswerRow, error) {
	periodDay := fmt.Sprintf("-%d days", period)
	query := `
	SELECT * FROM daily_log
	WHERE answeredAt >= datetime('now', ?);
	`

	response, err := conn.db.Query(query, periodDay)

	userAnswerRow := make([]UserAnswerRow, 0)

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
		userAnswerRow = append(userAnswerRow, row)
	}
	return userAnswerRow, err
}
