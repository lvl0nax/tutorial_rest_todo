package main

import (
	"log"

	"github.com/lvl0nax/tutorial_rest_todo/pkg/handler"
	"github.com/lvl0nax/tutorial_rest_todo/pkg/repository"
	"github.com/lvl0nax/tutorial_rest_todo/pkg/service"
	"github.com/spf13/viper"

	todo "github.com/lvl0nax/tutorial_rest_todo"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Config initializing error: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("http server run failed. Error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
