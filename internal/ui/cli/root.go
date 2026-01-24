/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tulashvili/MyDailyControlQuestions/internal/models"
	"github.com/tulashvili/MyDailyControlQuestions/internal/service"
)

var rootCmd = &cobra.Command{
	Use:   "MyDailyControlQuestions",
	Short: "Ежедневная саморефлексия путем ответа на важные для тебя вопросы по 5-ти бальной шкале",
	Long:  ``,
}

func Execute(conn *sql.DB) {
	db = conn

	statsCmd.Flags().IntVarP(&period, "period", "p", 1, "Stats over the period")
	statsCmd.MarkFlagRequired("stats")

	rootCmd.AddCommand(listQuestionsCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
