package auth

import (
	"fiber/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return err
	}
	return c.SendString(user.Username + user.Password)

	//return c.SendString("Login endpoint")
}

func Register(c *fiber.Ctx) error {

	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	userData := User{
		Username: user.Username,
		Password: user.Password,
	}

	db := database.DbConn()

	db.Create(&userData)

	//return c.SendString("User Registered Successfully")
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{"user": userData, "message": "User Registered Successfully"})
	//return c.SendString("Register endpoint")
}
