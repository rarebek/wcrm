CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT,
    image TEXT 
);


CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY, 
    worker_id UUID,
    product_id SERIAL,
    tax INT,
    discount INT,
    total_price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders_products(
    id SERIAL PRIMARY KEY,
    order_id UUID REFERENCES  orders(id),
    product_id INT
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(65),
    description VARCHAR(255),
    price INT,
    discount INT, 
    picture TEXT,
    category_id INT REFERENCES categories(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE owners (
    id UUID NOT NULL PRIMARY KEY,
    full_name VARCHAR(65) NOT NULL,
    company_name VARCHAR(65) NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    avatar TEXT,
    tax INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE workers (
    id UUID NOT NULL PRIMARY KEY,
    full_name VARCHAR(65) NOT NULL,
    login_key TEXT NOT NULL,
    password TEXT NOT NULL,
    owner_id UUID REFERENCES owners(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE geolocations (
    id SERIAL PRIMARY KEY,
	latitude FLOAT NOT NULL,
	longitude FLOAT NOT NULL,
    owner_id UUID REFERENCES owners(id)
);
