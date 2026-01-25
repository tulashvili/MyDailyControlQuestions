/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tulashvili/MyDailyControlQuestions/internal/usecase"
)

var listQuestionsCmd = &cobra.Command{
	Use:   "list_questions",
	Short: "Получить список текущих вопросов",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		questions := usecase.GetQuestionList()

		for _, question := range questions {
			fmt.Println("Категория:", question.Category)
			fmt.Println("Вопрос:", question.Text)
			fmt.Println()
		}
	},
}
