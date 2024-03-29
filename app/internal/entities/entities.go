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
	StoreProduct
	Data time.Time `json:"data"`
}

type InsertProduct struct {
	Product
	StoreAddress  string `json:"store_address"`
	CategoryTitle string `json:"category"`
	Amount        int    `json:"amount"`
}

type InsertCart struct {
	CustomerID int `json:"customer_id"`
	ProductID  int `json:"product_id"`
}
type OrderProduct struct {
	Product
	Amount int `json:"amount"`
}

type Order struct {
	ID       int            `json:"id"`
	Status   string         `json:"status"`
	Products []OrderProduct `json:"products"`
}
