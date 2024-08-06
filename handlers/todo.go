package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/maikkundev/start-daily-todo/database"
	"github.com/maikkundev/start-daily-todo/models"
)

// /todos
func TodosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTodos(w, r)
	case http.MethodPost:
		AddTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// /todo
func TodoHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	switch r.Method {
	case http.MethodGet:
		GetTodo(w, r, id)
	case http.MethodPut:
		UpdateTodo(w, r, id)
	case http.MethodDelete:
		DeleteTodo(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	database.Database.Find(&todos)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request, id string) {
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

func UpdateTodo(w http.ResponseWriter, r *http.Request, id string) {
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

func DeleteTodo(w http.ResponseWriter, r *http.Request, id string) {
	var todo models.Todo
	result := database.Database.Delete(&todo, id)

	if result.RowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
