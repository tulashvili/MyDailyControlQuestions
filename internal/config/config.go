package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/tulashvili/MyDailyControlQuestions/pkg/db"
)

type Config struct {
	DbPath db.DataSource
}

func NewConfig(loadDotEnv bool) (*Config, error) {
	if err := os.Setenv("TZ", "UTC"); err != nil {
		return nil, err
	}

	if loadDotEnv {
		if err := godotenv.Load(".env"); err == nil {
			slog.Info("loaded .env file")
		}
	}

	c := &Config{
		DbPath: db.LoadDatasource(),
	}

	if c.DbPath.SqlitePath == "" {
		return nil, fmt.Errorf("DB_PATH is empty")
	}

	return c, nil
}
