/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"database/sql"

	"github.com/spf13/cobra"
	repository "github.com/tulashvili/MyDailyControlQuestions/internal/repo"
	"github.com/tulashvili/MyDailyControlQuestions/pkg"
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
			pkg.FormatedPrintResult(row)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	statsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
