package repo

import "github.com/tulashvili/MyDailyControlQuestions/internal/domain"

type AnswerRepository interface {
	SaveAnswer(answer domain.UserAnswer) error
	GetAnswers(period int) ([]domain.UserAnswer, error)
}
