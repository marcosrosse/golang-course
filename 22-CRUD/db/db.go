package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver connection
)

// Conn open a connection with mysql
func Conn() (*sql.DB, error) {
	conn_string := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", conn_string)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
