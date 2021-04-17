package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
	db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/GREWords")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db
}
