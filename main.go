package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maikkundev/start-daily-todo/database"
	"github.com/maikkundev/start-daily-todo/handlers"
	"github.com/rs/cors"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to StartDailyToDo"))
	})

	mux.HandleFunc("/todos", handlers.TodosHandler)
	mux.HandleFunc("/todos/", handlers.TodoHandler)

	handler := cors.Default().Handler(mux)
	fmt.Printf("listening on port :3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
