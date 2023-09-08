CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
