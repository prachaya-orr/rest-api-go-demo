package main

import (
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBookById)
	app.Post("/books", createBook)
	app.Patch("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Listen(":8080")
}
