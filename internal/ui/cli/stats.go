/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"database/sql"
	"fmt"

	"github.com/spf13/cobra"
	repository "github.com/tulashvili/MyDailyControlQuestions/internal/repo"
)

var (
	db     *sql.DB
	period int
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Получить статистику ответов за Х дней",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := repository.SelectRows(db, period)
		if err != nil {
			panic(err)
		}

		for _, row := range rows {
			FormatedPrintResult(row)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}

func FormatedPrintResult(userAnswer repository.UserAnswerRow) {
	fmt.Println("---------------------")
	fmt.Println("ID:", userAnswer.ID)
	fmt.Println("Дата:", userAnswer.AnsweredAt)
	fmt.Println("Категория:", userAnswer.Category)
	fmt.Println("Вопрос:", userAnswer.Question)
	fmt.Println("Ответ:", userAnswer.Answer)
}
