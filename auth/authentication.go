package auth

import (
	"fiber/database"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(user User) []*ErrorResponse {

	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func Login(c *fiber.Ctx) error {

	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  err,
			"status": false})
	}

	errors := ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)

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

	//return c.JSON(userData.Username)
	status, msg := VerifyPassword(user.Password, userData.Password)

	if status {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		})

		token, err := claims.SignedString([]byte("secret"))

		if err != nil {
			return err
		}

		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Minute * 5),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":  true,
			"message": "User Logged-In Successfully",
			"user":    user,
			"token":   token})

	} else {

		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  msg,
			"status": false})

	}

	//return c.SendString(user.Username + user.Password)

	//return c.SendString("Login endpoint")
}

func UserInfo(c *fiber.Ctx) error {
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

	user := new(User)
	//var result User
	db.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func Register(c *fiber.Ctx) error {

	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  err,
			"status": false})
	}

	errors := ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

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

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{"message": "User Logged out Successfully"})

}
func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "Password is incorrect"
		check = false
	}

	return check, msg
}
