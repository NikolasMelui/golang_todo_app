package main

import (
	"os"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"

	todo "github.com/nikolasmelui/golang_todo_app"
	"github.com/nikolasmelui/golang_todo_app/pkg/handler"
	"github.com/nikolasmelui/golang_todo_app/pkg/repository"
	"github.com/nikolasmelui/golang_todo_app/pkg/service"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		DBName:   viper.GetString("db.db_name"),
		SSLMode:  viper.GetString("db.ssl_mode"),
		Password: viper.GetString("db.password"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialized db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("app.port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.Set("db.password", os.Getenv("DB_PASSWORD"))
	viper.Set("security.salt", os.Getenv("SECURITY_SALT"))
	viper.Set("security.signing_key", os.Getenv("SECURITY_SIGNING_KEY"))
	return viper.ReadInConfig()
}
