package domain

import "time"

type UserAnswer struct {
	ID               int
	QuestionCategory string
	QuestionText     string
	Answer           int
	AnsweredAt       *time.Time
}

// Init user answer
func CreateUserAnswer(question Question, answer int) UserAnswer {
	now := time.Now()

	return UserAnswer{
		QuestionCategory: string(question.Category),
		QuestionText:     question.Text,
		Answer:           answer,
		AnsweredAt:       &now,
	}
}