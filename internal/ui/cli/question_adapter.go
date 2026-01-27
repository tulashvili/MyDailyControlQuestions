package cli

import (
	"fmt"

	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
)

type QuestionAdapter struct{}

func (QuestionAdapter) ShowQuestion(question domain.Question) {
	fmt.Println("Категория:", question.Category)
	fmt.Println("Вопрос:", question.Text)
}

func (QuestionAdapter) ReadAnswer() (int, error) {
	var answer int
	_, err := fmt.Scanln(&answer)

	return answer, err
}

func (QuestionAdapter) ShowValidationError(err error) {
	fmt.Println("Ошибка:", err)
}
