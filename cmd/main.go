package main

import (
	"log"

	todo "github.com/nikolasmelui/golang_todo_app"
	"github.com/nikolasmelui/golang_todo_app/pkg/handler"
	"github.com/nikolasmelui/golang_todo_app/pkg/repository"
	"github.com/nikolasmelui/golang_todo_app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.AddConfigPath("config")
	return viper.ReadInConfig()
}
