package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(connString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
