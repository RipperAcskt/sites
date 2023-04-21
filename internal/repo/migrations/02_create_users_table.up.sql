CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    order_id VARCHAR(100) NOT NULL,
    state VARCHAR(30) NOT NULL,
    balance FLOAT(8) NOT NULL
);