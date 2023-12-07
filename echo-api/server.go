package main

import (
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
	middelware "github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "John", Age: 20},
	{ID: 2, Name: "Ohm", Age: 26},
}

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func getUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

type Err struct {
	Message string `json:"message"`
}

func createUserHandler(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	users = append(users, u)

	return c.JSON(http.StatusCreated, u)
}

func AuthMiddleware(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "1234" {
		return true, nil
	}
	return false, nil
}

func main() {
	e := echo.New()

	e.Use(middelware.Logger())
	e.Use(middelware.Recover())

	e.GET("/health", healthHandler)

	g := e.Group("/auth")
	g.Use(middelware.BasicAuth(AuthMiddleware))
	g.GET("/users", getUsersHandler)
	g.POST("/users", createUserHandler)

	log.Println("Server stared at http://localhost:2566")
	log.Fatal(e.Start(":2566"))
}
