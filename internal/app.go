package app

import (
	"database/sql"
	"log"

	"github.com/tulashvili/MyDailyControlQuestions/internal/config"
	"github.com/tulashvili/MyDailyControlQuestions/internal/ui/cli"

	"github.com/tulashvili/MyDailyControlQuestions/pkg/db/sqlitedb"
)

type App struct {
	DB     *sql.DB
	Config config.Config
}

func NewApp(conf config.Config) error {
	app := &App{
		Config: conf,
	}
	// Logic???
	// answers := usecase.AskQuestion(questions)

	// DB
	var err error
	app.DB, err = sqlitedb.InitDB(conf.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlitedb.CreateTable(app.DB); err != nil {
		log.Fatal(err)
	}

	// for _, answer := range answers {
	// 	if err := repository.InsertRow(app.DB, answer); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	cli.Execute(app.DB)
	return err
}
