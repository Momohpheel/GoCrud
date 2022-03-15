package book

import (
	"fiber/database"

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

	book := new(Book)

	db := database.DbConn()

	result := db.Find(&book)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  result.Error,
			"status": false})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "All the Books",
		"books":   book,
		"status":  true})

	//return c.SendString("Get All Books endpoint")
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book := new(Book)

	db := database.DbConn()

	result := db.Where("id = ?", id).Find(&book)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  result.Error,
			"status": false})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "One Book",
		"books":   book,
		"status":  true})

	//return c.SendString("Get One Book endpoint")
}

func AddBook(c *fiber.Ctx) error {

	book := new(Book)

	db := database.DbConn()

	if err := c.BodyParser(book); err != nil {
		return err
	}

	bookData := Book{
		Title:  book.Title,
		Author: book.Author,
		Isbn:   book.Isbn,
	}

	result := db.Create(bookData)

	if result.Error != nil {
		//return result.Error
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  result.Error,
			"status": false})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "One Book",
		"books":   book,
		"status":  true})

	//return c.SendString("Create New Book endpoint")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete One Book endpoint")
}
