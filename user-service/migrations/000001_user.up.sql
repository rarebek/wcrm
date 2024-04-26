CREATE TABLE owners (
    id UUID NOT NULL,
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
    id UUID NOT NULL,
    full_name VARCHAR(65) NOT NULL,
    login_key TEXT NOT NULL,
    password TEXT NOT NULL,
    owner_id UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE geolocations (
    id SERIAL PRIMARY KEY,
	latitude FLOAT NOT NULL,
	longitude FLOAT NOT NULL,
    owner_id UUID NOT NULL
);