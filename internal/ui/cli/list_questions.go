/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tulashvili/MyDailyControlQuestions/internal/service"
)

var listQuestionsCmd = &cobra.Command{
	Use:   "list_questions",
	Short: "Список текущих вопросов",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		questions := service.GetQuestions()

		for _, question := range questions {
			fmt.Println("Категория:", question.Category)
			fmt.Println("Вопрос:", question.Text)
			fmt.Println()
		}
	},
}
