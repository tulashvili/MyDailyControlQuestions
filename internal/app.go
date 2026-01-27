package app

import (
	"database/sql"
	"log"

	"github.com/tulashvili/MyDailyControlQuestions/internal/config"

	"github.com/tulashvili/MyDailyControlQuestions/pkg/db/sqlitedb"
)

type App struct {
	DB     *sql.DB
	Config config.Config
}

func NewApp(conf config.Config) (App, error) {
	app := &App{
		Config: conf,
	}

	// DB
	var err error
	app.DB, err = sqlitedb.InitDB(conf.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlitedb.CreateTable(app.DB); err != nil {
		log.Fatal(err)
	}
	return App{
		DB: app.DB,
	}, err
}
