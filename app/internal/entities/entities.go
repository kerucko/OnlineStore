package entities

type Customer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Password string `json:"password"`
}

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Shop        string `json:"shop"`
}

type Category struct {
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}
