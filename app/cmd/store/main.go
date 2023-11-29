package main

import (
	"OnlineStore/internal/config"
	"OnlineStore/internal/handlers"
	"OnlineStore/internal/storage/postgres"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
)

func main() {
	cfg := config.MustLoad()
	dbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	log.Println(dbPath)
	db, err := postgres.New(dbPath, cfg.Timeout)
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/product", handlers.GetProductHandler(db, cfg.Timeout))
	// пишем обработчики
}
