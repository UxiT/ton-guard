package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func (db *Database) Close() {
	err := db.DB.Close()
	if err != nil {
		panic(errors.Join(err, fmt.Errorf("failed to close database")))
	}
}

func NewDatabase(dbURL string) *Database {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	if err = db.Ping(); err != nil {
		panic("Failed to ping database: " + err.Error())
	}

	return &Database{db}
}
