package postgres

import (
	"OnlineStore/internal/entities"
	"OnlineStore/internal/storage"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	conn *pgx.Conn
}

func New(dbPath string, timeout time.Duration) (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := pgx.Connect(ctx, dbPath)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("Successful database connection")
	return &Storage{conn: conn}, nil
}

func (s *Storage) GetCustomerByID(id int, ctx context.Context) (entities.Customer, error) {
	request := "select customer_name, email, phone, address from customer where id = $1"
	var customer entities.Customer
	customer.ID = id
	row := s.conn.QueryRow(ctx, request, id)
	err := row.Scan(&customer.Name, &customer.Email, &customer.Phone, &customer.Address)
	if err != nil {
		return entities.Customer{}, storage.ErrNotExist
	}
	return customer, nil
}

func (s *Storage) GetProductByID(id int, ctx context.Context) (entities.Product, error) {
	request := `
	select p.title, p.description, p.price, p.photo_path, s.title
    	from product p 
    	    join store_product on p.id = store_product.product_id
    		join store on store_product.store_id = store.id
    		join seller s on store.seller_id = s.id
    	where p.id = $1
    `
	var product entities.Product
	product.ID = id
	row := s.conn.QueryRow(ctx, request, id)
	err := row.Scan(&product.Title, &product.Description, &product.Price, &product.PhotoPath, &product.Shop)
	if err != nil {
		return entities.Product{}, storage.ErrNotExist
	}
	return product, nil
}

func (s *Storage) GetAllProductFromCategory(categoryName string, ctx context.Context) (entities.Category, error) {
	var category entities.Category
	category.Name = categoryName
	request := "select p.id, p.title, p.price, p.photo_path from product p join category c on p.category_id = c.id where c.title = $1"
	rows, err := s.conn.Query(ctx, request, categoryName)
	if err != nil {
		return entities.Category{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var product entities.Product
		err = rows.Scan(&product.ID, &product.Title, &product.Price, &product.PhotoPath)
		if err != nil {
			return entities.Category{}, err
		}
		category.Products = append(category.Products, product)
	}
	if category.Products == nil {
		return entities.Category{}, storage.ErrNotExist
	}
	return category, nil
}

func (s *Storage) GetCustomerByEmail(email string, ctx context.Context) (entities.Customer, error) {
	request := `
		SELECT id, customer_name, email, phone, address
		FROM customer
		WHERE email = $1
	`
	var customer entities.Customer
	row := s.conn.QueryRow(ctx, request, email)
	err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.Address)
	if err != nil {
		return entities.Customer{}, storage.ErrNotExist
	}
	return customer, nil
}

func (s *Storage) GetAllSellersProducts(sellerID int, ctx context.Context) ([]entities.StoreProduct, error) {
	request := `
		select p.title, p.photo_path, sp.amount, s.address
		from product p 
		join store_product sp on p.id = sp.product_id
		join store s on sp.store_id = s.id
		join seller on s.seller_id = seller.id
		where seller.id = $1
		order by p.title
	`
	var products []entities.StoreProduct
	rows, err := s.conn.Query(ctx, request, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product entities.StoreProduct
		err = rows.Scan(&product.Title, &product.PhotoPath, &product.Amount, &product.StoreAddress)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
