package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer db.Close()
	//Query all
	stmt, err := db.Prepare("SELECT id, name, age FROM users where id = $1")
	if err != nil {
		log.Fatal("can't prepare query all users statement", err)
	}

	rowId := 1
	rows := stmt.QueryRow(rowId)

	var id, age int
	var name string

	err = rows.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal("can't scan row into variable", err)
	}

	fmt.Println("one row", id, name, age)
}
