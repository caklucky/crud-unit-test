package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = "kryptopos!"
	database string = "golang"
	host     string = "localhost"
	port     string = "3306"
)

var (
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
