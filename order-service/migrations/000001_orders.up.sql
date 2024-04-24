CREATE TABLE IF NOT EXISTS orders (
    id INT, 
    woker_id UUID,
    product_id INT,
    tax INT,
    discount INT,
    total_price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updeted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);