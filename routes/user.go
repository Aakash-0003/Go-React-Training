package routes

import (
	"github.com/Aakash-0003/Go-React-Training/controller"
	"github.com/gofiber/fiber/v2"
)

func RootSetup(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
}
