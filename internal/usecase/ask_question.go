package usecase

import (
	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
	"github.com/tulashvili/MyDailyControlQuestions/internal/repo"
)

type AskQuestionUC struct {
	repo repo.AnswerRepository
}

func NewAskQuestion(repo repo.AnswerRepository) *AskQuestionUC {
	return &AskQuestionUC{repo: repo}
}

func (aq *AskQuestionUC) Execute(questions []domain.Question, io QuestionIO) error {
	for _, q := range questions {
		for {
			io.ShowQuestion(q)

			answer, err := io.ReadAnswer()
			if err != nil {
				return err
			}

			if err := domain.ValidateAnswer(answer); err != nil {
				io.ShowValidationError(err)
			}

			userAnswer := domain.CreateUserAnswer(q, answer)

			if err := aq.repo.SaveAnswer(userAnswer); err != nil {
				return err
			}

			break
		}
	}
	return nil
}
