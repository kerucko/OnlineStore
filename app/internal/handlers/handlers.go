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

type Category struct {
	Name     string    `json:"name"`
	Products []Product `json:"products"`
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

func GetCategoryHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetCategoryHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		category := r.URL.Query().Get("category")
		log.Printf("%s %s", op, category)

		ps := Category{
			Name: "test",
			Products: []Product{
				{
					ID:          1,
					Name:        "Product 1",
					Price:       10,
					Description: "This is product 1",
					Shop:        "Shop A",
				},
				{
					ID:          2,
					Name:        "Product 2",
					Price:       20,
					Description: "This is product 2",
					Shop:        "Shop B",
				},
				{
					ID:          3,
					Name:        "Product 3",
					Price:       30,
					Description: "This is product 3",
					Shop:        "Shop C",
				},
				{
					ID:          4,
					Name:        "Product 4",
					Price:       40,
					Description: "This is product 4",
					Shop:        "Shop D",
				},
				{
					ID:          5,
					Name:        "Product 5",
					Price:       50,
					Description: "This is product 5",
					Shop:        "Shop E",
				},
				{
					ID:          6,
					Name:        "Product 6",
					Price:       60,
					Description: "This is product 6",
					Shop:        "Shop F",
				},
				{
					ID:          7,
					Name:        "Product 7",
					Price:       70,
					Description: "This is product 7",
					Shop:        "Shop G",
				},
				{
					ID:          8,
					Name:        "Product 8",
					Price:       80,
					Description: "This is product 8",
					Shop:        "Shop H",
				},
				{
					ID:          9,
					Name:        "Product 9",
					Price:       90,
					Description: "This is product 9",
					Shop:        "Shop I",
				},
				{
					ID:          10,
					Name:        "Product 10",
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
