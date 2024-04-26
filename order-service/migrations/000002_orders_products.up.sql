CREATE TABLE orders_products(
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES  orders(id),
    product_id INT
);