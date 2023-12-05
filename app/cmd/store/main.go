package main

import (
	"OnlineStore/internal/config"
	"OnlineStore/internal/handlers"
	"OnlineStore/internal/storage/postgres"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.MustLoad()
	dbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host, cfg.Postgres.Port, cfg.User, cfg.Password, cfg.DBName)
	log.Println(dbPath)
	db, err := postgres.New(dbPath, cfg.Timeout)
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/product", handlers.GetProductHandler(db, cfg.Timeout))
	router.Get("/show_all", handlers.GetCategoryHandler(db, cfg.Timeout))
	router.Get("/profile", handlers.GetCustomerProfileHandler(db, cfg.Timeout))
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))

}
