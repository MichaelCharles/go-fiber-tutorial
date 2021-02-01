package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/mcaubrey/go-fiber-tutorial/book"
	"github.com/mcaubrey/go-fiber-tutorial/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, 世界")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	// Book Endpoints
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened.")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database migrated...")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	app.Listen(3000)
}
