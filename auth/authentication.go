package auth

import (
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

	//return c.Send(c.Body())
	//return c.SendString("Login endpoint")
}

func Register(c *fiber.Ctx) error {

	return c.SendString("Register endpoint")
}
