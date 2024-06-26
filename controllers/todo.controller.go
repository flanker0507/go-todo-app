package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-todo-app/database"
	"go-todo-app/models"
	"go-todo-app/request"
	"log"
)

func CreateTodo(ctx *fiber.Ctx) error {
	todoReq := request.TodoUpdateeRequest{}

	// PARSE REQUEST BODY
	if errParse := ctx.BodyParser(&todoReq); errParse != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION DATA REQUEST
	validate := validator.New()
	if errValidate := validate.Struct(&todoReq); errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	// INSERT DATA KE DATABASE
	todo := models.Todo{}
	todo.Name = todoReq.Name
	todo.IsComplete = todoReq.IsComplete
	if todoReq.Note != "" {
		todo.Note = &todoReq.Note
	}

	if errDB := database.DB.Create(&todo).Error; errDB != nil {
		log.Println("todo.controller.go ==> CreateTodo :: ", errDB)
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return ctx.Status(201).JSON(fiber.Map{
		"message": "todo crated successfully",
		"data":    todo,
	})
}

func GetAllTodo(ctx *fiber.Ctx) error {
	todos := []models.Todo{}
	if err := database.DB.Find(&todos).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "data transmited",
		"data":    todos,
	})
}

func GetTodoById(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")
	todo := models.Todo{}

	if err := database.DB.First(&todo, "id = ?", &todoId).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "todo not found",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "data transmited",
		"data":    todo,
	})
}

func UpdateTodoById(ctx *fiber.Ctx) error {
	todoReq := request.TodoUpdateRequest{}

	// PARSE REQUEST BODY
	if errParse := ctx.BodyParser(&todoReq); errParse != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "fail to parsing data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATION DATA REQUEST
	validate := validator.New()
	if errValidate := validate.Struct(&todoReq); errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	todoId := ctx.Params("id")
	todo := models.Todo{}

	if err := database.DB.First(&todo, "id = ?", &todoId).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "todo not found",
		})
	}
	todo.Name = todoReq.Name
	todo.Note = &todoReq.Note
	todo.IsComplete = todoReq.IsComplete

	if errSave := database.DB.Save(&todo).Error; errSave != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "todo updated",
		"data":    todo,
	})
}

func DeleteTodoById(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")
	todo := models.Todo{}

	if err := database.DB.First(&todo, "id = ?", &todoId).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "todo not found",
		})
	}

	if errDel := database.DB.Delete(&todo).Error; errDel != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "todo deleted",
	})
}
