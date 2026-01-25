package repo

import "github.com/tulashvili/MyDailyControlQuestions/internal/domain"

type AnswerRepository interface {
	Save(answer domain.UserAnswer) error
	List(period int) ([]domain.UserAnswer, error)
}
