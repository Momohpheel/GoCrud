package main

import (
	"fiber/auth"
	"fiber/book"
	"fiber/middleware"
	"fmt"
	"log"
	"os"

	"fiber/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func setupRoute(app *fiber.App) {
	group := app.Group("api/v1")
	group.Post("/login", auth.Login)
	group.Post("/register", auth.Register)
	group.Get("/user", auth.UserInfo)
	group.Post("/logout", auth.Logout)

	//group.Use()

	group.Get("/books", book.GetAllBooks)
	group.Get("/books/:id", book.GetBook)
	group.Post("/book", book.AddBook)
	group.Delete("/books/:id", book.DeleteBook)
}
func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	result := fiber.New()

	result.Use(middleware.Authenticate(result))

	result.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	db := database.DbConn()

	fmt.Println(os.Getenv("DB_DATABASE"))

	db.AutoMigrate(&auth.User{})
	db.AutoMigrate(&book.Book{})

	setupRoute(result)

	log.Fatal(result.Listen(":3000"))

}
