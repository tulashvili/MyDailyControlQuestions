package db

import (
	"fmt"
	"os"
)

type Datasource struct {
	SqlitePath string
}

func LoadDatasource() Datasource {
	wd, _ := os.Getwd()
	fmt.Println("CWD:", wd)
	fmt.Println("DB_PATH:", os.Getenv("DB_PATH"))
	return Datasource{
		SqlitePath: os.Getenv("DB_PATH"),
	}
}
