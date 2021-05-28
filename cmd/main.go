package main

import (
	"log"

	todo "github.com/nikolasmelui/golang_todo_app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
