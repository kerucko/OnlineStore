-- Количество товаров, пришедших на склад в определенный промежуток времени !store_id
SELECT store.id, SUM(delivery.amount)
FROM delivery
JOIN store ON delivery.store_id = store.id AND store.id = 2
WHERE delivery.date BETWEEN '2023-08-20' AND '2023-08-25'
GROUP BY store.id;

-- Имена пользователей, которые купили хотя бы три товара из категории 'Техника'
SELECT op.order_id, c.customer_name
FROM customer c
JOIN orders o ON o.customer_id = c.id
JOIN order_product op ON op.order_id = o.id
JOIN product p ON op.product_id = p.id
JOIN category ON p.category_id = category.id
WHERE category.title = 'Техника'
GROUP BY op.order_id, c.customer_name
HAVING SUM(op.amount) > 3;

-- Имя пользователя, который совершил заказ на максимальную сумму (подзапрос)
SELECT c.customer_name, c.address, SUM(p.price*op.amount) AS price
FROM customer c
JOIN orders o ON c.id = o.customer_id
JOIN order_product op on o.id = op.order_id
JOIN product p on op.product_id = p.id
GROUP BY op.order_id, c.customer_name, c.address
HAVING SUM(p.price*op.amount) = (
  SELECT SUM(p.price * op.amount) AS Стоимость
  FROM order_product op
  JOIN product p ON op.product_id = p.id
  GROUP BY order_id 
  ORDER BY Стоимость DESC
  LIMIT 1
);

-- Вывести все поставки на все склады определенного продавца !seller_id
SELECT d.product_id, d.store_id, d.date, d.amount
FROM delivery d
JOIN store s ON d.store_id = s.id
WHERE s.seller_id = 1
ORDER BY d.store_id ASC, d.date DESC;

-- Топ 10 самых популярных товаров
SELECT p.title, SUM(op.amount) AS total
FROM product p
JOIN order_product op ON p.id = op.product_id
GROUP BY p.title
ORDER BY total DESC
LIMIT 10;

-- Вывести всех пользователей, у которых хотя бы один заказ доставлен.
SELECT customer_name, email, phone
FROM customer 
join orders on customer.id = orders.customer_id
WHERE orders.customer_id = ANY(
	SELECT customer_id 
	FROM orders 
	WHERE status = 'Доставлен'
);

-- Изменить все статусы заказов 'Получен' на 'Доставлен'
UPDATE orders 
SET status = 'Доставлен'
WHERE status = 'Получен';

-- Все товары из категорий 'Женская одежда и обувь' и 'Мужская одежда и обувь'
SELECT p.title, p.price, p.url_photo, c.title AS category_title, p.description
FROM product p
JOIN category c ON p.category_id = c.id
WHERE c.title = 'Женская одежда и обувь'
UNION
SELECT p.title, p.price, p.url_photo, c.title AS category_title, p.description
FROM product p
JOIN category c ON p.category_id = c.id
WHERE c.title = 'Мужская одежда и обувь';

-- Все пользователи, купившие товары и в категории "Техника", и в категории "Книги"
SELECT customer_name, email, phone
FROM customer
JOIN orders ON customer.id = orders.customer_id
JOIN order_product ON orders.id = order_product.order_id
JOIN product ON order_product.product_id = product.id
JOIN category ON product.category_id = category.id
WHERE category.title = 'Техника'
INTERSECT
SELECT customer_name, email, phone
FROM customer
JOIN orders ON customer.id = orders.customer_id
JOIN order_product ON orders.id = order_product.order_id
JOIN product ON order_product.product_id = product.id
JOIN category ON product.category_id = category.id
WHERE category.title = 'Книги';

-- Все продавцы, которые продают товары из категории "Спорт", но не продают товары из категории "Техника"
SELECT seller.title
FROM seller
JOIN store ON seller.id = store.seller_id
JOIN store_product ON store.id = store_product.store_id
JOIN product ON store_product.product_id = product.id
JOIN category ON product.category_id = category.id
WHERE category.title = 'Спорт'
EXCEPT
SELECT seller.title
FROM seller
JOIN store ON seller.id = store.seller_id
JOIN store_product ON store.id = store_product.store_id
JOIN product ON store_product.product_id = product.id
JOIN category ON product.category_id = category.id
WHERE category.title = 'Техника';

-- Вывести всех пользователей, у которых хотя бы один заказ был оформлен хотя бы в одной из категорий Красота, Спорт, Здоровье и количество товаров было равно 3.
SELECT customer_name, email, phone, address 
FROM customer 
WHERE EXISTS (
	SELECT * 
	FROM orders
	JOIN order_product ON order_product.order_id = orders.id
	JOIN product ON order_product.product_id = product.id
	JOIN category ON product.category_id = category.id
	WHERE category.title IN ('Красота', 'Спорт', 'Здоровье') AND orders.customer_id = customer.id AND order_product.amount = 3
)
ORDER BY customer_name;

-- Товары и комментарий, в котором будет указана цена выше или ниже средней цены, а также максимальная и минимальная цена
SELECT p.title, p.price, 
CASE
  WHEN p.price = (SELECT MAX(price) FROM product) THEN 'Самая большая цена'
  WHEN p.price = (SELECT MIN(price) FROM product) THEN 'Самая маленькая цена'
  WHEN p.price > (SELECT AVG(price) FROM product) THEN 'Цена выше средней'
  WHEN p.price < (SELECT AVG(price) FROM product) THEN 'Цена ниже средней'
  ELSE 'Цена равна средней'
END AS price_comment
FROM product p
ORDER BY p.price;

-- Вывести всех пользователей, у которых хотя бы один заказ доставлен.
SELECT customer_name, email, phone
FROM customer 
join orders on customer.id = orders.customer_id
WHERE orders.customer_id = SOME(
	SELECT customer_id 
	FROM orders 
	WHERE status = 'Доставлен'
);

-- Товары, цена которых выше средней цены всех категорий, не включая товары категории "Техника"
SELECT p.title, p.price
FROM product p
JOIN category c ON p.category_id = c.id
WHERE c.title <> 'Техника' AND p.price > ALL (
  SELECT AVG(p2.price)
  FROM category c2
  JOIN product p2 ON c2.id = p2.category_id
  WHERE c2.title <> 'Техника'
  GROUP BY c2.title
)
ORDER BY p.price;