package main

import (
	backend_trainee_assignment_2023 "backend-trainee-assignment-2023"
	"backend-trainee-assignment-2023/pkg/handlers"
	"backend-trainee-assignment-2023/repository"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("failed initializing configs: %v\n", err.Error())
	}

	db, err := repository.ConnectToDb(repository.Config{
		User:     viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Name:     viper.GetString("db.name"),
	})
	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err.Error())
	}

	handler := handlers.InitRoutes()
	server := new(backend_trainee_assignment_2023.Server)
	if err := server.Run(viper.GetString("port"), handler); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("error occured while running http server: %v\n", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}
