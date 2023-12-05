package handlers

import (
	"OnlineStore/internal/entities"
	"OnlineStore/internal/storage"
	"OnlineStore/internal/storage/postgres"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetProductHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetProductHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		productID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Printf("%s %s", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		product, err := db.GetProductByID(productID, ctx)
		if errors.Is(err, storage.ErrNotExist) {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s sending reply %v", op, product)
	}
}

func GetCategoryHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetCategoryHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		category := r.URL.Query().Get("category")
		log.Printf("%s %s", op, category)

		ps := entities.Category{
			Name: "test",
			Products: []entities.Product{
				{
					ID:          1,
					Title:       "Product 1",
					Price:       10,
					Description: "This is product 1",
					Shop:        "Shop A",
				},
				{
					ID:          2,
					Title:       "Product 2",
					Price:       20,
					Description: "This is product 2",
					Shop:        "Shop B",
				},
				{
					ID:          3,
					Title:       "Product 3",
					Price:       30,
					Description: "This is product 3",
					Shop:        "Shop C",
				},
				{
					ID:          4,
					Title:       "Product 4",
					Price:       40,
					Description: "This is product 4",
					Shop:        "Shop D",
				},
				{
					ID:          5,
					Title:       "Product 5",
					Price:       50,
					Description: "This is product 5",
					Shop:        "Shop E",
				},
				{
					ID:          6,
					Title:       "Product 6",
					Price:       60,
					Description: "This is product 6",
					Shop:        "Shop F",
				},
				{
					ID:          7,
					Title:       "Product 7",
					Price:       70,
					Description: "This is product 7",
					Shop:        "Shop G",
				},
				{
					ID:          8,
					Title:       "Product 8",
					Price:       80,
					Description: "This is product 8",
					Shop:        "Shop H",
				},
				{
					ID:          9,
					Title:       "Product 9",
					Price:       90,
					Description: "This is product 9",
					Shop:        "Shop I",
				},
				{
					ID:          10,
					Title:       "Product 10",
					Price:       100,
					Description: "This is product 10",
					Shop:        "Shop J",
				},
			},
		}

		jsonBytes, err := json.Marshal(ps)
		if err != nil {
			log.Printf("%s %s", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, string(jsonBytes))
	}
}

func GetCustomerProfileHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetCustomerProfileHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		customerID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		customer, err := db.GetCustomerByID(customerID, ctx)
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
