package db

import (
	"os"
)

type DataSource struct {
	SqlitePath string
}

func LoadDatasource() DataSource {
	return DataSource{
		SqlitePath: os.Getenv("DB_PATH"),
	}
}
