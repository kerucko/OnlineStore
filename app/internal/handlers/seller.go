package handlers

import (
	"OnlineStore/internal/entities"
	"OnlineStore/internal/storage/postgres"
	"context"
	"encoding/json"
	"io"
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

func GetSellerStoresHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetSellerStoresHandler:"
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
		stores, err := db.GetSellerStores(sellerID, ctx)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(stores)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s sending reply %v", op, stores)
	}
}

func GetSellerDeliveriesHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetSellerDeliveriesHandler:"
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
		stores, err := db.GetSellerDeliveries(sellerID, ctx)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(stores)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s sending reply", op)
	}
}

func NewProductHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "NewProductHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		sellerID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Printf("%s %s", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("%s: ReadAll %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var object entities.InsertProduct
		err = json.Unmarshal(body, &object)
		if err != nil {
			log.Printf("%s: Unmarshal %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println(object)
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		err = db.AddNewProduct(ctx, object, sellerID)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
