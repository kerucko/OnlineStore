-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE category (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE product (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price BIGINT NOT NULL,
    url_photo VARCHAR(255) NOT NULL,
    category_id BIGINT NOT NULL REFERENCES category(id)
);

CREATE TABLE customer (
    id BIGSERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL
);

CREATE TABLE cart (
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES customer(id)
);

CREATE TABLE cart_product (
    cart_id BIGINT NOT NULL REFERENCES cart(id),
    product_id BIGINT NOT NULL REFERENCES product(id),
    amount BIGINT NOT NULL
);

CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL REFERENCES customer(id),
    date TIMESTAMP NOT NULL,
    status VARCHAR(255) NOT NULL
);

CREATE TABLE order_product (
    order_id BIGINT NOT NULL REFERENCES orders(id),
    product_id BIGINT NOT NULL REFERENCES product(id),
    amount BIGINT NOT NULL
);

CREATE TABLE seller (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE store (
    id BIGSERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    seller_id BIGINT NOT NULL REFERENCES seller(id)
);

CREATE TABLE store_product (
    store_id BIGINT NOT NULL REFERENCES store(id),
    product_id BIGINT NOT NULL REFERENCES product(id),
    amount BIGINT NOT NULL
);

CREATE TABLE delivery (
    store_id BIGINT NOT NULL REFERENCES store(id),
    product_id BIGINT NOT NULL REFERENCES product(id),
    amount BIGINT NOT NULL,
    date TIMESTAMP NOT NULL
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE order_product;
DROP TABLE cart_product;
DROP TABLE store_product;
DROP TABLE delivery;
DROP TABLE orders;
DROP TABLE cart;
DROP TABLE store;
DROP TABLE customer;
DROP TABLE seller;
DROP TABLE product;
DROP TABLE category;
