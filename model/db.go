package model

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func EstablishConnection() (*sqlx.DB, error) {
	user := os.Getenv("DB_USERNAME")
	if user == "" {
		user = "root"
	}

	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		pass = "pass"
	}

	host := os.Getenv("DB_HOSTNAME")
	if host == "" {
		host = "localhost"
	}

	dbname := os.Getenv("DB_DATABASE")
	if dbname == "" {
		dbname = "linq"
	}

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, dbname)+"?parseTime=True&loc=Asia%2FTokyo&charset=utf8mb4")
	return db, err
}
