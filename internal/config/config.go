package config

import (
	"os"
)

type Config struct {
	DbPath string
}

func LoadConfig() Config {
	return Config{
		DbPath: os.Getenv("DB_PATH"),
	}
}
