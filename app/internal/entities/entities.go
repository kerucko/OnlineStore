package entities

import "time"

type Customer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	PhotoPath   string `json:"photo_path"`
	Shop        string `json:"shop"`
}

type Category struct {
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type StoreProduct struct {
	Title        string `json:"title"`
	Amount       int    `json:"amount"`
	StoreAddress string `json:"store_address"`
	PhotoPath    string `json:"photo_path"`
}

type Store struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
}

type Delivery struct {
	StoreID  int            `json:"store_id"`
	Data     time.Time      `json:"data"`
	Products []StoreProduct `json:"products"`
}

type InsertProduct struct {
	Product
	StoreAddress  string `json:"store_address"`
	CategoryTitle string `json:"category"`
	Amount        int    `json:"amount"`
}
