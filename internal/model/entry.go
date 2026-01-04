package model

import "time"

type Entry struct {
	ID               int
	QuestionCategory string
	QuestionText     string
	Answer           int
	CompletedAt      *time.Time
}
