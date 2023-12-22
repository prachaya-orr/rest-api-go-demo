package main

import (
	"fmt"
	"log"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	html "github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	},
	)

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBookById)
	app.Post("/books", createBook)
	app.Patch("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)
	app.Get("/test-html", testHTML)

	app.Get("/config", getEnv)

	app.Listen(":8080")
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	fmt.Println(file, err)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	c.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File uploaded successfully!")
}

func testHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
		"Name":  "Go Learner",
	})
}

func getEnv(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("SECRET"),
	})
}
