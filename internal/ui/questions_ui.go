package ui

import (
	"fmt"

	"github.com/tulashvili/MyDailyControlQuestions/internal/models"
	"github.com/tulashvili/MyDailyControlQuestions/internal/service"
)

// Get question list with category
func GetQuestions(questions []models.Question) {
	for _, question := range questions {
		fmt.Println("Категория:", question.Category)
		fmt.Println("Вопрос:", question.Text)
	}
}

// Ask a question to a client
func AskQuestion(questions []models.Question) []models.UserAnswer {
	var result []models.UserAnswer

	for _, question := range questions {
		fmt.Println("Категория:", question.Category)
		fmt.Println("Вопрос:", question.Text)

		var answer int
		fmt.Scanln(&answer)

		err := service.ValidateAnswer(answer)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Введи значение заново")
			fmt.Scanln(&answer)
		} else {
			userAnswer := service.CreateUserAnswer(question, answer)
			result = append(result, userAnswer)
		}
	}
	return result
}

func PrintResult(result []models.UserAnswer) {
	for _, v := range result {
		fmt.Println(v)
	}
}
