package pkg

import (
	"fmt"

	repository "github.com/tulashvili/MyDailyControlQuestions/internal/repo"
)

// Formated result from SelectRow
func FormatedPrintResult(userAnswer repository.UserAnswerRow) {
	fmt.Println("---------------------")
	fmt.Println("ID:", userAnswer.ID)
	fmt.Println("Дата:", userAnswer.AnsweredAt)
	fmt.Println("Категория:", userAnswer.Category)
	fmt.Println("Вопрос:", userAnswer.Question)
	fmt.Println("Ответ:", userAnswer.Answer)
}
