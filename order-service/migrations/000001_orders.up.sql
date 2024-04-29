CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY, 
    worker_id UUID,
    product_id INT,
    tax INT,
    discount INT,
    total_price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);