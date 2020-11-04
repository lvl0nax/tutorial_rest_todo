package main

import (
	"log"

	"github.com/lvl0nax/tutorial_rest_todo/pkg/handler"

	todo "github.com/lvl0nax/tutorial_rest_todo"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("http server run failed. Error: %s", err.Error())
	}
}
