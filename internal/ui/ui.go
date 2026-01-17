package ui

import (
	"database/sql"
	"fmt"

	"github.com/tulashvili/MyDailyControlQuestions/internal/models"
	"github.com/tulashvili/MyDailyControlQuestions/internal/service"
	"github.com/tulashvili/MyDailyControlQuestions/internal/repo"
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

// Show data over period
func ShowDataOverPeriod(conn *sql.DB) error {
	period := 2
	rows, err := repository.SelectRows(conn, period)
	if err != nil {
		return err
	}

	for _, row := range rows {
		formatedPrintResult(row)
	}
	return nil
}

// Formated result from SelectRow
func formatedPrintResult(userAnswer repository.UserAnswerRow) {
	fmt.Println("---------------------")
	fmt.Println("ID:", userAnswer.ID)
	fmt.Println("Дата:", userAnswer.AnsweredAt)
	fmt.Println("Категория:", userAnswer.Category)
	fmt.Println("Вопрос:", userAnswer.Question)
	fmt.Println("Ответ:", userAnswer.Answer)
}
