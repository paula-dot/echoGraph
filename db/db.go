package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(user, password, dbname string) error {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
    }
	return DB.Ping()
