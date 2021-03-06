package routes

import (
	"github.com/Aakash-0003/Go-React-Training/controller"
	"github.com/gofiber/fiber/v2"
)

func RootSetup(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
	app.Get("/user", controller.User)
	app.Post("/logout", controller.Logout)
	app.Post("/role", controller.AdminRoleUpdate)
	app.Post("/admindelete", controller.AdminDelete)
	app.Get("/clockin", controller.ClockIn)
}
