package main

import (
	"github.com/gofiber/fiber/v2"
	"go-todo-app/routes"
)

func main() {

	app := fiber.New()
	//init route
	routes.InitRoute(app)

	app.Listen(":8080")

}
