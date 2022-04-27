package main

import (
	"github.com/Aakash-0003/Go-React-Training/database"
	"github.com/Aakash-0003/Go-React-Training/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.CreateConnection()
	app := fiber.New()
	routes.RootSetup(app)

	app.Listen(":8000")
}
