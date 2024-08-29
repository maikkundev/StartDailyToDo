package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/maikkundev/start-daily-todo/internal/database"
	"github.com/maikkundev/start-daily-todo/internal/models"
)

func (s server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/todos", GetTodos)
	r.Get("/todo/{id}", GetTodo)
	r.Post("/todo", AddTodo)
	r.Patch("/todo/{id}", UpdateTodo)
	r.Delete("/todo/{id}", DeleteTodo)

	return r
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	database.Database.Find(&todos)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var todo models.Todo
	result := database.Database.First(&todo, id)

	if result.RowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.Database.Create(&todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.Database.Model(&todo).Where("id = ?", id).Updates(map[string]interface{}{
		"IsDone": todo.IsDone,
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var todo models.Todo
	result := database.Database.Delete(&todo, id)

	if result.RowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
