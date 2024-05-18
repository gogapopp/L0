CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    data JSONB
);