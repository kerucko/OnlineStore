package postgres

import (
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
