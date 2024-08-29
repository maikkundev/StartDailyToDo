package main

import (
	"fmt"

	"github.com/maikkundev/start-daily-todo/internal/server"
)

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Cannot start server: %s", err.Error()))
	}
}
