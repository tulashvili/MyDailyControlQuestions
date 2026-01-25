package usecase

// import (
// 	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
// )

// type AskQuestionUseCase struct {
// 	input Input
// }

// // Ask a question to a client
// func (uc *AskQuestionUseCase) Run(questions []domain.Question) ([]domain.UserAnswer, error) {
// 	var result []domain.UserAnswer
// 	var answer int

// 	for _, question := range questions {
// 		if err := domain.ValidateAnswer(answer); err != nil {
// 			return nil, err
// 		}
// 		userAnswer := domain.CreateUserAnswer(question, answer)
// 		result = append(result, userAnswer)
// 	}
// 	return result, err
// }
