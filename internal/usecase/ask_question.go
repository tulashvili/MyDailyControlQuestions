package usecase

import (
	"time"

	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
)

// Init user answer
func CreateUserAnswer(question domain.Question, answer int) domain.UserAnswer {
	now := time.Now()

	return domain.UserAnswer{
		QuestionCategory: string(question.Category),
		QuestionText:     question.Text,
		Answer:           answer,
		AnsweredAt:       &now,
	}
}
