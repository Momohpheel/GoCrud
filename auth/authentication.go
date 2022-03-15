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
	p := new(Person)

        if err := c.BodyParser(p); err != nil {
            return err
        }

        log.Println(p.Name) // john
        log.Println(p.Pass) // doe
	//return c.Send(c.Body())
	//return c.SendString("Login endpoint")
}

func Register(c *fiber.Ctx) error {

	return c.SendString("Register endpoint")
}
