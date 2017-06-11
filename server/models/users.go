package models

import (
	"database/sql"
	"errors"
	"golang.org/x/net/context"
)

type User struct {
	Username string
}

func AllUsers(ctx context.Context) ([]*User, error) {
	db, ok := ctx.Value("db").(*sql.DB)
	if !ok {
		return nil, errors.New("models: could not get database connection pool from context")
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*User, 0)
	for rows.Next() {
		bk := new(User)
		err := rows.Scan(&bk.Username)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}