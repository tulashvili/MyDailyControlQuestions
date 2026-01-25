/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tulashvili/MyDailyControlQuestions/internal/models"
	"github.com/tulashvili/MyDailyControlQuestions/internal/service"
)

// askQuestionsCmd represents the askQuestions command
var askQuestionsCmd = &cobra.Command{
	Use:   "ask",
	Short: "Запустить опрос",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// askQuestion()

	},
}

// Ask a question to a client
func askQuestion(questions []models.Question) []models.UserAnswer {
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
