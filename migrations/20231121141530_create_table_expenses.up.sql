CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    description VARCHAR(255),
    incoming DECIMAL(10,2),
    outgoing DECIMAL(10,2),
    balance DECIMAL(10,2),
    date DATE,
    category_id INTEGER,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
