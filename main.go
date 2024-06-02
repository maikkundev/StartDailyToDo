package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maikkundev/start-daily-todo/database"
	"github.com/maikkundev/start-daily-todo/handlers"
)

func main() {
	app := fiber.New()

	err := database.Connect()
	if err != nil {
		return
	}

	app.Get("/todos", handlers.GetTodos)
	app.Get("todos/:id", handlers.GetTodo)
	app.Post("/todos", handlers.AddTodo)
	app.Put("/todos/:id", handlers.UpdateTodo)
	app.Delete("todos/:id", handlers.DeleteTodo)

	log.Fatal(app.Listen(":3000"))
}
