CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY, 
    worker_id UUID,
    product_id SERIAL,
    tax INT,
    discount INT,
    total_price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(65),
    description VARCHAR(255),
    price INT,
    discount INT, -- it should be under than 100 and higer than 0
    picture TEXT,
    category_id INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE orders_products(
    id SERIAL,
    order_id SERIAL REFERENCES  orders(id),
    product_id SERIAL REFERENCES products(id)
);

