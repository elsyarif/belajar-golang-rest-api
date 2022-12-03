package mysql

import (
	"github.com/jmoiron/sqlx"
	"time"
)

func NewConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/golang_db?parseTime=true")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db, nil
}
