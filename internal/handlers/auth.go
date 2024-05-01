package handlers

import (
	"OnlineStore/internal/storage"
	"OnlineStore/internal/storage/postgres"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

func SignInHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "SignInHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		email := r.URL.Query().Get("email")
		// password := r.URL.Query().Get("password")

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		customer, err := db.GetCustomerByEmail(email, ctx)
		switch {
		case errors.Is(err, storage.ErrNotExist):
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusNotFound)
			return
		case err == nil:
			log.Printf("%s Success", op)
		default:
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(customer)
		if err != nil {
			log.Printf("%s encode: %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s sending reply %v", op, customer)
	}
}
