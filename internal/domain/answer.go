package domain

import "time"

type UserAnswer struct {
	ID               int
	QuestionCategory string
	QuestionText     string
	Answer           int
	AnsweredAt       *time.Time
}
