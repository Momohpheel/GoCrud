package main

import (
	"fiber/auth"
	"fiber/book"
	"log"

	"fiber/database"

	"github.com/gofiber/fiber/v2"
)

func setupRoute(app *fiber.App) {
	app.Post("api/v1/login", auth.Login)
	app.Post("api/v1/register", auth.Login)

	//app.Use()

	app.Get("/api/v1/books", book.GetAllBooks)
	app.Get("/api/v1/books/:id", book.GetBook)
	app.Post("/api/v1/book", book.AddBook)
	app.Delete("/api/v1/books/:id", book.DeleteBook)
}
func main() {
	result := fiber.New()
	db := database.DbConn()

	db.AutoMigrate(&auth.User{})
	db.AutoMigrate(&book.Book{})

	setupRoute(result)

	log.Fatal(result.Listen(":3000"))

}
