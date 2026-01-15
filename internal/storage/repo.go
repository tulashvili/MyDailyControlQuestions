package storage

import (
	"database/sql"
	"fmt"

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
func SelectRows(conn *sql.DB, period int) ([]UserAnswerRow, error) {
	periodDay := fmt.Sprintf("-%d days", period)
	query := `
	SELECT * FROM daily_log
	WHERE answeredAt >= datetime('now', ?);
	`

	response, err := conn.Query(query, periodDay)

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
		formatedPrintResult(row)
	}
	return userAnswerRow, err
}

// Formated result from SelectRow
func formatedPrintResult(userAnswer UserAnswerRow) {
	fmt.Println("---------------------")
	fmt.Println("ID:", userAnswer.ID)
	fmt.Println("Дата:", userAnswer.AnsweredAt)
	fmt.Println("Категория:", userAnswer.Category)
	fmt.Println("Вопрос:", userAnswer.Question)
	fmt.Println("Ответ:", userAnswer.Answer)
}
