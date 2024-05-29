CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY,
    owner_id TEXT,
    title VARCHAR(65),
    description VARCHAR(255),
    price INT,
    discount INT, 
    picture TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY,
    owner_id TEXT,
    name TEXT,
    image TEXT, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories_products (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    product_id  UUID REFERENCES products(id) ON DELETE CASCADE,
    category_id UUID REFERENCES categories(id) ON DELETE CASCADE, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    worker_id UUID,
    products JSONB,
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


CREATE TABLE owners (
    id UUID NOT NULL PRIMARY KEY,
    full_name VARCHAR(65) NOT NULL,
    company_name VARCHAR(65) NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    avatar TEXT,
    tax INT,
    refresh_token TEXT,
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
