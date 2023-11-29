package handlers

import (
	"OnlineStore/internal/storage/postgres"
	"net/http"
	"time"
)

func GetProductHandler(db postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetProductHandler: "
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
