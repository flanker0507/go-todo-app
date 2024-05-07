package routes

import "github.com/gofiber/fiber/v2"

func v1Route(app *fiber.App) {

	v1 := app.Group("v1")

	todo := v1.Group("/todo")

	todo.Post("/", func(ctx *fiber.Ctx) error {

		//create one todo

		return nil
	})
	todo.Get("/", func(ctx *fiber.Ctx) error {

		//get list of todo

		return nil
	})

	todo.Get("/:id", func(ctx *fiber.Ctx) error {

		//get todo by id

		return nil
	})

	todo.Patch("/:id", func(ctx *fiber.Ctx) error {

		//update todo by id

		return nil
	})

	todo.Delete("/:id", func(ctx *fiber.Ctx) error {

		//delete todo by id

		return nil
	})

}
