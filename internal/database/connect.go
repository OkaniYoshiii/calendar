package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const MAX_IDLE_TIME = time.Minute * 3
const MAX_OPEN_CONNS = 5

func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(MAX_IDLE_TIME)
	db.SetMaxOpenConns(MAX_OPEN_CONNS)
	db.SetMaxIdleConns(MAX_OPEN_CONNS)

	return db, nil
}
