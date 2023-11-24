package main

import "OnlineStore/internal/config"

func main() {
	cfg := config.MustLoad()
	_ = cfg
	// подключаем логер

	// подлючаем базу данных

	// пишем обработчики
}
