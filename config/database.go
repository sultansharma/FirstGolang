package config

import (
	_ "database/sql"

	"log"
   "github.com/jmoiron/sqlx"
   _ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func GetMySQLDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "swati:123456@(localhost:3306)/merubazar?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
