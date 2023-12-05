package postgres

import (
	"OnlineStore/internal/entities"
	"OnlineStore/internal/storage"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
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
