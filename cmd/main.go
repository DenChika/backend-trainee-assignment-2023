package main

import (
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

	router := handlers.InitRoutes()
	server := http.Server{
		Addr:    ":" + viper.GetString("port"),
		Handler: router,
	}
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("failed to start server: %v\n", err.Error())
	}
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}
