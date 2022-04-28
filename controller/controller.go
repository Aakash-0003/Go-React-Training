package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/Aakash-0003/Go-React-Training/database"
	"github.com/Aakash-0003/Go-React-Training/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const Secret = "98123f84fc6e29380c1a9bf8628cf63b7ac071ac5b6ae45bb821a7cd39030a08"

func Register(c *fiber.Ctx) error {

	//fmt.Printf("%v \n%t", c)

	var data map[string]string
	err := c.BodyParser(&data)
	database.ErrorHandling(err)
	fmt.Println("data: ", data)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 15)
	//convert data into db model
	user := models.User{
		Username: data["name"],
		Email:    data["email"],
		Password: hashedPassword,
	}

	//insert user data(data converted to DB  model user ) into database
	inserted := database.AddUser(&user)
	return c.JSON(inserted)
}

//login
func Login(c *fiber.Ctx) error {

	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data: ", data)
	userMail := data["email"]
	var user models.User
	database.FindUserByMail(&userMail, &user)
	if user.Id.IsZero() {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	//return c.JSON(user)

	//JWT IMPLEMENTATION USING CLAIMS

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{

		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(Secret))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}
	fmt.Println("token generated successfully: ", token)
	cookie := fiber.Cookie{
		Name:     "jwt",
		Expires:  time.Now().Add(time.Hour * 24),
		Value:    token,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message code": c.Response().StatusCode(),
		"message":      "Login Successfull \n Token Generated",
		"token":        token,
	})

}

//parse, verify ,validate the signature from cookie of user who's logged in
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt") // to grab the cookie in the request body from current login user

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("token : ", token)
		return []byte(Secret), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	fmt.Println("data: ", claims.Issuer)
	database.FindUserByMail(&claims.Issuer, &user)

	return c.JSON(user)
}
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "successfully logged out",
	})
}
