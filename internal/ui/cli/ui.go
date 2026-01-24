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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "MyDailyControlQuestions",
	Short: "Ежедневная саморефлексия путем ответа на важные для тебя вопросы по 5-ти бальной шкале",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(conn *sql.DB, ) {
	db = conn

	statsCmd.Flags().IntVarP(&period, "period", "p", 1, "Stats over the period")
	statsCmd.MarkFlagRequired("stats")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.MyDailyControlQuestions.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
