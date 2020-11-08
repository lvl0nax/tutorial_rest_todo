// docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/lvl0nax/tutorial_rest_todo/pkg/handler"
	"github.com/lvl0nax/tutorial_rest_todo/pkg/repository"
	"github.com/lvl0nax/tutorial_rest_todo/pkg/service"

	todo "github.com/lvl0nax/tutorial_rest_todo"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Config initializing error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed loading ENV variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.ssl_mode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("DB connection failed. Error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
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
