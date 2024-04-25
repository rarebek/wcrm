CREATE TABLE orders_products(
    id SERIAL,
    order_id SERIAL REFERENCES  orders(id),
    product_id SERIAL REFERENCES products(id)
);