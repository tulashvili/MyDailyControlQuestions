package usecase

import (
	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
	"github.com/tulashvili/MyDailyControlQuestions/internal/repo"
)

type GetStatAnswer struct {
	repo repo.AnswerRepository
}

func NewGetStat(repo repo.AnswerRepository) *GetStatAnswer {
	return &GetStatAnswer{repo: repo}
}

func (gs *GetStatAnswer) Execute(period int) ([]domain.UserAnswer, error) {
	return gs.repo.GetAnswers(period)
}
