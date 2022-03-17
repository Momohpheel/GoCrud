package middleware

import (
	"fiber/auth"
	"fiber/database"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error":  "Unauthorized",
			"status": false})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	db := database.DbConn()

	user := new(auth.User)
	//var result User
	db.Where("id = ?", claims.Issuer).First(&user)

	return c.Next()

}
