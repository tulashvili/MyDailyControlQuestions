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

// func GetStatAnswers(period int) []domain.UserAnswer {

// 	rows, err := repository.SelectRows(db, period)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, row := range rows {
// 		formatedPrintResult(row)
// 	}
// }

// func formatedPrintResult(userAnswer repository.UserAnswerRow) {
// 	fmt.Println("---------------------")
// 	fmt.Println("ID:", userAnswer.ID)
// 	fmt.Println("Дата:", userAnswer.AnsweredAt)
// 	fmt.Println("Категория:", userAnswer.Category)
// 	fmt.Println("Вопрос:", userAnswer.Question)
// 	fmt.Println("Ответ:", userAnswer.Answer)
// }
