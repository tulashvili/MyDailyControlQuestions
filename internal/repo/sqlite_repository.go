package repo

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
)

type SQLiteAnswerRepository struct {
	db *sql.DB
}

func NewSQLiteAnswerRepository(db *sql.DB) *SQLiteAnswerRepository {
	return &SQLiteAnswerRepository{db: db}
}

// Save answers to questions
func (r *SQLiteAnswerRepository) SaveAnswer(answer domain.UserAnswer) error {
	query := `
        INSERT INTO daily_log (category, question, answer, answeredAt)
        VALUES (?, ?, ?, ?)
    `
	_, err := r.db.Exec(
		query,
		answer.QuestionCategory,
		answer.QuestionText,
		answer.Answer,
		answer.AnsweredAt.UTC().Format(time.RFC3339Nano),
	)
	return err
}

// Get data over some period
func (r *SQLiteAnswerRepository) GetAnswers(period int) ([]domain.UserAnswer, error) {
	periodDay := fmt.Sprintf("-%d days", period)
	query := `
        SELECT category, question, answer, answeredAt
        FROM daily_log
        WHERE answeredAt >= datetime('now', ?);
    `

	rows, err := r.db.Query(query, periodDay)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.UserAnswer
	var answeredAtStr sql.NullString

	for rows.Next() {
		var ua domain.UserAnswer

		if err := rows.Scan(
			&ua.QuestionCategory,
			&ua.QuestionText,
			&ua.Answer,
			&answeredAtStr,
		); err != nil {
			return nil, err
		}

		if answeredAtStr.Valid {
			t, err := time.Parse(time.RFC3339Nano, answeredAtStr.String)
			if err != nil {
				return nil, err
			}
			ua.AnsweredAt = &t
		}

		result = append(result, ua)
	}

	return result, nil
}
