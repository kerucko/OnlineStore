package handlers

import (
	"OnlineStore/internal/storage/postgres"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetSellersProductsHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetSellersProducts:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		sellerID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Printf("%s %s", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		products, err := db.GetAllSellersProducts(sellerID, ctx)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(products)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s sending reply %v", op, products)
	}
}
