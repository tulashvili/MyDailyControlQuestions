/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"github.com/spf13/cobra"
	"github.com/tulashvili/MyDailyControlQuestions/internal/domain"
	"github.com/tulashvili/MyDailyControlQuestions/internal/repo"
	"github.com/tulashvili/MyDailyControlQuestions/internal/usecase"
)

// askQuestionsCmd represents the askQuestions command
var askQuestionsCmd = &cobra.Command{
	Use:   "ask",
	Short: "Запустить опрос",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		repo := repo.NewSQLiteAnswerRepository(db)
		ask := usecase.NewAskQuestion(repo)
		io := QuestionAdapter{}
		
		if err := ask.Execute(domain.QUESTIONS, io); err != nil {
			panic(err)
		}
	},
}
