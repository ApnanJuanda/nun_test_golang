package config

import (
	"api/nun_test/helper"
	"database/sql"
	"time"
)

func GetMyConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/nun_db?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)
	return db
}
