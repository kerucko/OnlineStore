package main

import (
	"OnlineStore/internal/config"
	"OnlineStore/internal/handlers"
	"OnlineStore/internal/storage/postgres"
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	db, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("public"))
	router.Handle("/*", fs)

	router.Get("/product", handlers.GetProductHandler(db, cfg.Timeout))
	router.Get("/show_all", handlers.GetCategoryHandler(db, cfg.Timeout))
	router.Get("/profile", handlers.GetCustomerProfileHandler(db, cfg.Timeout))

	router.Get("/signin", handlers.SignInHandler(db, cfg.Timeout))
	router.Get("/cart", handlers.GetCartHandler(db, cfg.Timeout))
	router.Post("/cart", handlers.AddProductToCartHandler(db, cfg.Timeout))
	router.Get("/all_orders", handlers.GetAllOrdersHandler(db, cfg.Timeout))
	router.Get("/order", handlers.GetOrderHandler(db, cfg.Timeout))
	router.Get("/delete_from_cart", handlers.DeleteFromCartHandler(db, cfg.Timeout))

	router.Post("/buy", handlers.BuyHandler(db, cfg.Timeout))

	router.Get("/seller/product", handlers.GetSellersProductsHandler(db, cfg.Timeout))
	router.Get("/seller/store", handlers.GetSellerStoresHandler(db, cfg.Timeout))
	router.Get("/seller/delivery", handlers.GetSellerDeliveriesHandler(db, cfg.Timeout))

	router.Post("/seller/product", handlers.NewProductHandler(db, cfg.Timeout))
	router.Post("/seller/store", handlers.NewSellerStoreHandler(db, cfg.Timeout))
	log.Fatal(http.ListenAndServe("[::]:"+cfg.Server.Port, router))
}
