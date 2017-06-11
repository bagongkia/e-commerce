package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func OpenDB() (*DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/ecommerce")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}