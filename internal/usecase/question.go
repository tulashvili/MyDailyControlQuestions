package usecase

import "github.com/tulashvili/MyDailyControlQuestions/internal/domain"

type QuestionIO interface {
	ShowQuestion(domain.Question)
	ReadAnswer() (int, error)
	ShowValidationError(error)
}