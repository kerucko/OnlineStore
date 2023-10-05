CREATE TABLE category (
    id BIGSERIAL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE product (
    id BIGSERIAL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price BIGINT NOT NULL,
    url_photo VARCHAR(255) NOT NULL,
    category_id BIGINT NOT NULL
);

CREATE TABLE customer (
    id BIGSERIAL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL
);

CREATE TABLE cart (
    id BIGSERIAL,
    customer_id BIGINT NOT NULL
);

CREATE TABLE cart_product (
    cart_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    amount BIGINT NOT NULL
);

CREATE TABLE order (
    id BIGSERIAL,
    customer_id BIGINT NOT NULL,
    date TIMESTAMP NOT NULL,
    status VARCHAR(255) NOT NULL
);

CREATE TABLE order_product (
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    amount BIGINT NOT NULL
);

CREATE TABLE seller (
    id BIGSERIAL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE store (
    id BIGSERIAL,
    address VARCHAR(255) NOT NULL,
    seller_id BIGINT NOT NULL
);

CREATE TABLE store_product (
    store_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    amount BIGINT NOT NULL
);

CREATE TABLE delivery (
    prodict_id BIGINT NOT NULL,
    store_id BIGINT NOT NULL,
    date TIMESTAMP NOT NULL,
    amount BIGINT NOT NULL
);