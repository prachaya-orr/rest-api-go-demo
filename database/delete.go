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

	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1;")

	if err != nil {
		log.Fatal("can't prepare delete statemnet", err)
	}

	if _, err := stmt.Exec(1); err != nil {
		log.Fatal("can't execute delete statement", err)
	}

	fmt.Println("delete success")
}
