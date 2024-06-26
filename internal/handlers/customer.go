package handlers

import (
	"OnlineStore/internal/entities"
	"OnlineStore/internal/storage"
	"OnlineStore/internal/storage/postgres"
	"context"
	"encoding/json"
	"errors"
	"io"
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

		categoryName := r.URL.Query().Get("category")
		log.Printf("%s %s", op, categoryName)

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		var (
			category entities.Category
			err      error
		)

		defer cancel()
		if categoryName == "hits" {
			category, err = db.GetBestSellers(ctx)
		} else {
			category, err = db.GetAllProductFromCategory(categoryName, ctx)
		}
		switch {
		case errors.Is(err, storage.ErrNotExist):
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusNotFound)
			return
		case err == nil:
			log.Printf("%s success", op)
		default:
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(category)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s sending reply %v", op, category)
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

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
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

func GetCartHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetBasketHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		customerID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		cart, err := db.GetCart(ctx, customerID)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(cart)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s sending reply %v", op, cart)
	}
}

func AddProductToCartHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "AddProductToCartHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("%s: ReadAll %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var object entities.InsertCart
		err = json.Unmarshal(body, &object)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		err = db.AddProductToCart(ctx, object.ProductID, object.CustomerID)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s success", op)
	}
}

func GetAllOrdersHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetAllOrdersHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		customerID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		orders, err := db.GetAllOrders(ctx, customerID)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(orders)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s sending reply %v", op, orders)
	}
}

func GetOrderHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "GetOrderHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		orderID, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		order, err := db.GetOrder(ctx, orderID)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(order)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s sending reply %v", op, order)
	}
}

func DeleteFromCartHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "DeleteFromCartHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		customerID, err := strconv.Atoi(r.URL.Query().Get("customer_id"))
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		productID, err := strconv.Atoi(r.URL.Query().Get("product_id"))
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		err = db.DeleteFromCart(ctx, customerID, productID)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ok := true
		err = json.NewEncoder(w).Encode(ok)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s Success", op)
	}
}

func BuyHandler(db *postgres.Storage, timeout time.Duration) http.HandlerFunc {
	op := "BuyHandler:"
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")

		customerID, err := strconv.Atoi(r.URL.Query().Get("customer_id"))
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		err = db.BuyCart(ctx, customerID)
		if err != nil {
			log.Printf("%s %v", op, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("%s Success", op)
	}
}
