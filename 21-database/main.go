package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn_string := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", conn_string)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("The connection is opened!")
	}

	lines, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(lines)
	}
	defer lines.Close()

}
