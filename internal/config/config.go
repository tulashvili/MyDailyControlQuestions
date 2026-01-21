package config

import (
	"github.com/tulashvili/MyDailyControlQuestions/pkg/db"
)

type Config struct {
	DbPath db.DataSource
}

// func LoadConfig() Config {
// 	return Config{
// 		DbPath: os.Getenv("DB_PATH"),
// 	}
// }
