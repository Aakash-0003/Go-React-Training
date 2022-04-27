package controller

import (
	"fmt"

	"github.com/Aakash-0003/Go-React-Training/database"
	"github.com/Aakash-0003/Go-React-Training/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

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
	database.ErrorHandling(err)
	fmt.Println("data: ", data)
	userMail := data["email"]
	var user models.User
	database.FindUser(&userMail, &user)
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
	return c.JSON(user)

}
