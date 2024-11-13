CREATE TABLE IF NOT EXISTS orders(
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL, FOREIGN KEY (user_id) REFERENCES users(id),
	status VARCHAR(50) DEFAULT 'pending', -- Example: pending, completed, etc.
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_orders_user_id ON orders (user_id);
