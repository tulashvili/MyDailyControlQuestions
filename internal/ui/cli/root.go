/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"database/sql"

	"github.com/spf13/cobra"
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
	rootCmd.AddCommand(askQuestionsCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
