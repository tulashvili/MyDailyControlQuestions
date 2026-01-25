package repo

import "github.com/tulashvili/MyDailyControlQuestions/internal/domain"

type AnswerRepository interface {
	SaveAnswers(answer domain.UserAnswer) error
	GetAnswers(period int) ([]domain.UserAnswer, error)
}