package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB(address string) {
	db, err := sql.Open("mysql", address)
	if err != nil {
		log.Fatal("MySQl param error:" + err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("MySQL connect error:" + err.Error())
	}
	DB = db
}

func CloseDB() {
	if DB != nil {
		defer DB.Close()
	}
}
