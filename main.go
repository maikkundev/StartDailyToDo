package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/maikkundev/start-daily-todo/database"
	"github.com/maikkundev/start-daily-todo/handlers"
)

func main() {
	var app = fiber.New()
	app.Use(cors.New())

	var err = database.Connect()
	if err != nil {
		return
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to StartDailyToDo")
	})

	app.Get("/todos", handlers.GetTodos)
	app.Get("todos/:id", handlers.GetTodo)
	app.Post("/todos", handlers.AddTodo)
	app.Put("/todos/:id", handlers.UpdateTodo)
	app.Delete("todos/:id", handlers.DeleteTodo)

	log.Fatal(app.Listen(":3000"))
}
