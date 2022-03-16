package auth

import (
	"fiber/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  err,
			"status": false})
	}

	userData := User{
		Username: user.Username,
		Password: user.Password,
	}

	db := database.DbConn()

	//var result User
	result := db.Where("username = ?", user.Username).First(&user)

	//db.QueryFields("SELECT * FROM users WHERE username=? AND password=?", user.Username, user.Password)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  result.Error,
			"status": false})
	}

	status, msg := VerifyPassword(userData.Password, user.Password)

	if status {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "User Logged-In Successfully",
			"user":    user})

	} else {

		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  msg,
			"status": false})

	}

	//return c.SendString(user.Username + user.Password)

	//return c.SendString("Login endpoint")
}

func Register(c *fiber.Ctx) error {

	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  err,
			"status": false})
	}

	userData := User{
		Username: user.Username,
		Password: hashPassword(user.Password),
	}

	db := database.DbConn()

	db.Create(&userData)

	//return c.SendString("User Registered Successfully")
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{"user": userData, "message": "User Registered Successfully"})
	//return c.SendString("Register endpoint")
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "Password is incorrect"
		check = false
	}

	return check, msg
}
