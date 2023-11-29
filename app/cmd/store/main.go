package main

import (
	"OnlineStore/internal/config"
	"OnlineStore/internal/storage/postgres"
	"fmt"
	"log"
)

func main() {
	cfg := config.MustLoad()
	dbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err := postgres.New(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	_ = db
	// пишем обработчики
}
