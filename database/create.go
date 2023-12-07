package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("main init")
}

func main() {
	// ohm.Say()
	url := "postgres://jzwmkewr:Mc7eIhYXLL84_7pAP3TqUf4-czLf-MVF@rain.db.elephantsql.com/jzwmkewr"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, age INT);
	`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("failed to create table", err)
	}

	log.Println("connected to database")
}
