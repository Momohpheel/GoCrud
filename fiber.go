package main

import (
	"fiber/auth"
	"fiber/book"
	"fmt"
	"log"
	"os"

	"fiber/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoute(app *fiber.App) {
	app.Post("api/v1/login", auth.Login)
	app.Post("api/v1/register", auth.Register)

	//app.Use()

	app.Get("/api/v1/books", book.GetAllBooks)
	app.Get("/api/v1/books/:id", book.GetBook)
	app.Post("/api/v1/book", book.AddBook)
	app.Delete("/api/v1/books/:id", book.DeleteBook)
}
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	result := fiber.New()
	db := database.DbConn()

	fmt.Println(os.Getenv("DB_DATABASE"))

	db.AutoMigrate(&auth.User{})
	db.AutoMigrate(&book.Book{})

	setupRoute(result)

	log.Fatal(result.Listen(":3000"))

}
