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
	latitude VARCHAR(65) NOT NULL,
	longitude VARCHAR(65) NOT NULL,
    owner_id UUID REFERENCES owners(id)
);