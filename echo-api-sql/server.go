package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	echo "github.com/labstack/echo/v4"
	middelware "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type Err struct {
	Message string `json:"message"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, age INT);`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table users", err)
	}

	e := echo.New()

	e.Use(middelware.Logger())
	e.Use(middelware.Recover())

	e.GET("/health", healthHandler)
	e.GET("/users", getUsersHandler)
	e.POST("/users", createUserHandler)

	g := e.Group("/auth")
	g.Use(middelware.BasicAuth(AuthMiddleware))
	g.GET("/users", getUsersHandler)
	g.POST("/users", createUserHandler)

	log.Println("Server stared at http://localhost:2566")
	log.Fatal(e.Start(":2566"))
}

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func getUsersHandler(c echo.Context) error {
	stmt, err := db.Prepare("SELECT id,name,age FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	users := []User{}

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}

func createUserHandler(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	row := db.QueryRow("INSERT INTO users (name,age) values ($1,$2) RETURNING id", u.Name, u.Age)

	err = row.Scan(&u.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, u)
}

func AuthMiddleware(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "1234" {
		return true, nil
	}
	return false, nil
}
