package models

import (
	"time"
)

type UserAnswer struct {
	ID               int
	QuestionCategory string
	QuestionText     string
	Answer           int
	AnsweredAt       *time.Time
}

type Question struct {
	Category
	Text     string
}

type Category string
