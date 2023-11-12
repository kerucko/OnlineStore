-- Все пользователи и их заказанные продукты (GROUP BY)
SELECT customer.customer_name, product.title
FROM customer
JOIN cart ON customer.id = cart.customer_id
JOIN cart_product ON cart.id = cart_product.cart_id
JOIN product ON cart_product.product_id = product.id
GROUP BY customer.customer_name, product.title;

-- Количество товаров, пришедших на склад в определенный промежуток времени (BETWEEN, GROUP BY)
SELECT store.id, SUM(delivery.amount)
FROM delivery
JOIN store ON delivery.store_id = store.id AND store.id = 2
WHERE delivery.date BETWEEN '2023-08-20' AND '2023-08-25'
GROUP BY store.id;

-- Имена пользователей, которые купили хотя бы три товара из категории 'Книги' (SUM)
SELECT c.customer_name, c.address, p.title
FROM customer c
JOIN orders o ON o.customer_id = c.id
JOIN order_product op ON op.order_id = o.id
JOIN product p ON op.product_id = p.id
JOIN category ON p.category_id = category.id
WHERE category.title = 'Техника'
GROUP BY c.customer_name, c.address, p.title, op.amount
HAVING SUM(op.amount) > 3;

-- Имя пользователя, который совершил заказ на максимальную сумму (подзапрос)

-- Вывести все поставки на все склады определенного продавца

-- Топ 10 самых популярных товаров (ORDER BY)

-- Сколько человек добавили в корзину конкретный товар

-- Все товары, которые заказывали оптом (ANY)