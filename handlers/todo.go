package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maikkundev/start-daily-todo/database"
	"github.com/maikkundev/start-daily-todo/models"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo

	database.Database.Find(&todos)
	return c.Status(200).JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo

	result := database.Database.Find(&todo, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(result)
}

func AddTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	database.Database.Create(&todo)
	return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	id := c.Params("id")

	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Where("id = ?", id).Updates(&todo)
	return c.Status(200).JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo

	result := database.Database.Delete(&todo, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
