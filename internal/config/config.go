package config

import "github.com/tulashvili/MyDailyControlQuestions/pkg/db"

type Config struct {
	DS db.Datasource
}

func NewConfig() (Config, error) {
	cfg := Config{}
	return cfg, nil
}
