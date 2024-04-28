package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"

	"OnlineStore/internal/config"
	"OnlineStore/internal/entities"
	"OnlineStore/internal/storage"
)

type Storage struct {
	conn *pgx.Conn
}

func New(ctx context.Context, cfg config.Postgres) (*Storage, error) {
	dbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	deadline := time.After(cfg.Timeout)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			conn, err := pgx.Connect(ctx, dbPath)
			if err != nil {
				continue
			}
			if err = conn.Ping(ctx); err != nil {
				continue
			}
			log.Println("Successful database connection")
			return &Storage{conn: conn}, nil

		case <-deadline:
			return nil, fmt.Errorf("unable to connect to database")
		}
	}
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
	request := `
		select p.id, p.title, p.price, p.photo_path 
		from product p 
		join category c on p.category_id = c.id 
		join store_product st on p.id = st.product_id
		join store on st.store_id = store.id
		where c.title = $1
	`
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

func (s *Storage) GetBestSellers(ctx context.Context) (entities.Category, error) {
	request := `
		SELECT p.id, p.title, p.price, p.photo_path
		FROM product p
		JOIN order_product op ON p.id = op.product_id
		GROUP BY p.id, p.title
		ORDER BY SUM(op.amount) DESC
		LIMIT 10;
	`
	var category entities.Category
	category.Name = "Бестселлеры"
	rows, err := s.conn.Query(ctx, request)
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

func (s *Storage) GetSellerStores(sellerID int, ctx context.Context) ([]entities.Store, error) {
	request := `
		select store.id, store.address
		from store
		join seller on store.seller_id = seller.id
		where seller.id = $1
	`
	var stores []entities.Store
	rows, err := s.conn.Query(ctx, request, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var store entities.Store
		err = rows.Scan(&store.ID, &store.Address)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}
	return stores, nil
}

func (s *Storage) GetSellerDeliveries(sellerID int, ctx context.Context) ([]entities.Delivery, error) {
	request := `
		select p.title, p.photo_path, d.date, d.amount, s.address
		from product p 
		join delivery d on p.id = d.product_id
		join store s on d.store_id = s.id
		join seller on s.seller_id = seller.id
		where seller.id = $1
		order by s.id, d.date
	`
	var deliveries []entities.Delivery
	rows, err := s.conn.Query(ctx, request, sellerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var delivery entities.Delivery
		err = rows.Scan(&delivery.Title, &delivery.PhotoPath, &delivery.Data, &delivery.Amount, &delivery.StoreAddress)
		if err != nil {
			return nil, err
		}

		deliveries = append(deliveries, delivery)
	}
	return deliveries, nil
}

func (s *Storage) getCategoryID(ctx context.Context, category string) (int, error) {
	getCategorySQL := "select id from category where title = $1"
	var id int
	row := s.conn.QueryRow(ctx, getCategorySQL, category)
	err := row.Scan(&id)
	if err != nil {
		return 0, storage.ErrNotExist
	}
	return id, nil
}

func (s *Storage) getStoreID(ctx context.Context, sellerID int, address string) (int, error) {
	getCategorySQL := `
		select store.id 
		from store 
		join seller on store.seller_id = seller.id
		where address = $1 and seller.id = $2
	`
	var id int
	row := s.conn.QueryRow(ctx, getCategorySQL, address, sellerID)
	err := row.Scan(&id)
	if err != nil {
		return 0, storage.ErrNotExist
	}
	return id, nil
}

func (s *Storage) AddNewProduct(ctx context.Context, object entities.InsertProduct, sellerID int) error {
	categoryID, err := s.getCategoryID(ctx, object.CategoryTitle)
	if err != nil {
		return err
	}
	log.Println(categoryID)
	insertProductSQL := `
		insert into product(title, description, price, photo_path, category_id) 
		values ($1, $2, $3, $4, $5) RETURNING id
	`
	var productID int
	err = s.conn.QueryRow(ctx, insertProductSQL, object.Title, object.Description, object.Price, object.PhotoPath, categoryID).Scan(&productID)
	if err != nil {
		return err
	}
	storeID, err := s.getStoreID(ctx, sellerID, object.StoreAddress)
	if err != nil {
		return nil
	}
	log.Println(storeID)
	insertStoreProductSQL := `
		insert into store_product(store_id, product_id, amount)
		values($1, $2, $3)
	`

	_, err = s.conn.Exec(ctx, insertStoreProductSQL, storeID, productID, object.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetCart(ctx context.Context, customerID int) ([]entities.OrderProduct, error) {
	request := `
		SELECT p.id, p.title, p.price, p.photo_path, se.title, SUM(cp.amount)
		FROM product p
		JOIN cart_product cp ON p.id = cp.product_id
		JOIN cart c ON cp.cart_id = c.id AND c.customer_id = $1
		JOIN store_product sp ON p.id = sp.product_id
		JOIN store s ON sp.store_id = s.id
		JOIN seller se ON s.seller_id = se.id
		GROUP BY p.id, se.title
	`

	var cart []entities.OrderProduct
	rows, err := s.conn.Query(ctx, request, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entities.OrderProduct
		err = rows.Scan(&product.ID, &product.Title, &product.Price, &product.PhotoPath, &product.Shop, &product.Amount)
		if err != nil {
			return nil, err
		}
		cart = append(cart, product)
	}

	return cart, nil
}

func (s *Storage) AddNewStore(ctx context.Context, sellerID int, address string) error {
	insertSQL := "insert into store(address, seller_id) values($1, $2)"
	_, err := s.conn.Exec(ctx, insertSQL, address, sellerID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) getCartID(ctx context.Context, customerID int) (int, error) {
	request := "select cart.id from cart join customer on cart.customer_id = customer.id where customer.id=$1"
	var id int
	row := s.conn.QueryRow(ctx, request, customerID)
	err := row.Scan(&id)
	if err != nil {
		return 0, storage.ErrNotExist
	}
	return id, nil
}

func (s *Storage) AddProductToCart(ctx context.Context, productID int, customerID int) error {
	cartID, err := s.getCartID(ctx, customerID)
	if err != nil {
		return err
	}
	request := "insert into cart_product(cart_id, product_id, amount) values($1, $2, 1)"
	_, err = s.conn.Exec(ctx, request, cartID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetOrder(ctx context.Context, id int) (entities.Order, error) {
	request := `
		SELECT p.id, p.title, p.price, p.photo_path, se.title, op.amount
		FROM product p
		JOIN store_product sp ON p.id = sp.product_id
		JOIN store s ON sp.store_id = s.id
		JOIN seller se ON s.seller_id = se.id
		JOIN order_product op ON p.id = op.product_id
		JOIN orders o ON op.order_id = o.id AND o.id = $1
	`
	var order entities.Order
	order.ID = id
	rows, err := s.conn.Query(ctx, request, id)
	if err != nil {
		return entities.Order{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entities.OrderProduct
		err = rows.Scan(&product.ID, &product.Title, &product.Price, &product.PhotoPath, &product.Shop, &product.Amount)
		if err != nil {
			return entities.Order{}, err
		}
		order.Products = append(order.Products, product)
	}

	return order, nil
}

func (s *Storage) GetAllOrders(ctx context.Context, customerID int) ([]entities.Order, error) {
	request := `
		SELECT o.id, o.status
		FROM orders o
		JOIN customer c ON o.customer_id = c.id AND c.id = $1
	`

	var orders []entities.Order
	rows, err := s.conn.Query(ctx, request, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order entities.Order
		err = rows.Scan(&order.ID, &order.Status)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (s *Storage) DeleteFromCart(ctx context.Context, customerID int, productID int) error {
	deleteSQL := `
	DELETE FROM cart_product
	WHERE cart_id = (select id from cart where customer_id = $1) and product_id = $2`
	_, err := s.conn.Exec(ctx, deleteSQL, customerID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) BuyCart(ctx context.Context, customerID int) error {
	cart, err := s.GetCart(ctx, customerID)
	if err != nil {
		return err
	}
	insertOrderRequest := `
		INSERT INTO orders(customer_id, status, date) 
		VALUES($1, $2, NOW()) 
		RETURNING id
	`

	var orderID int
	err = s.conn.QueryRow(ctx, insertOrderRequest, customerID, "В сборке").Scan(&orderID)
	if err != nil {
		return err
	}

	for _, product := range cart {
		insertOrderProductRequest := `
			INSERT INTO order_product(order_id, product_id, amount) 
			VALUES($1, $2, $3)
		`
		_, err = s.conn.Exec(ctx, insertOrderProductRequest, orderID, product.ID, product.Amount)
		if err != nil {
			return err
		}
	}

	deleteCartRequest := `
		DELETE FROM cart_product
		WHERE cart_product.cart_id = $1
	`

	_, err = s.conn.Exec(ctx, deleteCartRequest, customerID)
	if err != nil {
		return err
	}
	log.Printf("Order %d created, customerID: %d", orderID, customerID)
	return nil
}
