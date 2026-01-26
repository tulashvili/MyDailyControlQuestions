/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
	"github.com/tulashvili/MyDailyControlQuestions/internal/repo"
	"github.com/tulashvili/MyDailyControlQuestions/internal/usecase"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Получить статистику ответов за Х дней",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		repo := repo.NewSQLiteAnswerRepository(db)
		stats, err := usecase.NewGetStat(repo).Execute(period)
		if err != nil {
			log.Fatal(err)
		}
		for _, answer := range stats {
			formatedPrintResult(answer)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}

func formatedPrintResult(userAnswer domain.UserAnswer) {
	fmt.Println("---------------------")
	fmt.Println("ID:", userAnswer.ID)
	fmt.Println("Дата:", userAnswer.AnsweredAt)
	fmt.Println("Категория:", userAnswer.QuestionCategory)
	fmt.Println("Вопрос:", userAnswer.QuestionText)
	fmt.Println("Ответ:", userAnswer.Answer)
}
