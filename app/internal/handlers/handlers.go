package handlers

import (
	"OnlineStore/internal/storage/postgres"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Shop        string `json:"shop"`
}

func GetProductHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetProductHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// ctx, cancel := context.WithTimeout(r.Context(), timeout)
		id := r.URL.Query().Get("id")
		_, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("%s %s", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("%s %s", op, id)

		p := Product{
			ID:          1,
			Name:        "test",
			Price:       1,
			Description: "test",
			Shop:        "test",
		}
		jsonBytes, err := json.Marshal(p)
		if err != nil {
			log.Printf("%s %s", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, string(jsonBytes))
	}
}
