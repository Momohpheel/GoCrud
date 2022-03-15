package book

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Isbn   uint   `json:"isbn"`
}

func GetAllBooks(c *fiber.Ctx) error {
	return c.SendString("Get All Books endpoint")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Get One Book endpoint")
}

func AddBook(c *fiber.Ctx) error {
	return c.SendString("Create New Book endpoint")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete One Book endpoint")
}
