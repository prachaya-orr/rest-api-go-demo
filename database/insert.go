package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer db.Close()

	// อยากได้ของให้ใช้ QueryRow ไม่อยากได้ ให้ใช้ Exec
	row := db.QueryRow("Insert into users (name, age) values ($1, $2) RETURNING id", "prachaya", 26)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("failed to insert row", err)
	}
	log.Println("inserted todo success id : ", id)

}
