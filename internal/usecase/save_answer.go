package usecase

import (
	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
	"github.com/tulashvili/MyDailyControlQuestions/internal/repo"
)

type SaveAnswer struct {
	repo repo.AnswerRepository
}

func NewSaveAnswers(repo repo.AnswerRepository) *SaveAnswer {
	return &SaveAnswer{repo: repo}
}

func (sa *SaveAnswer) Execute(answers domain.UserAnswer) error {
	return sa.repo.SaveAnswers(answers)
}
