CREATE TABLE IF NOT EXISTS orders (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL, FOREIGN KEY (user_id) REFERENCES users(id),
	delivery_time INTEGER NOT NULL, -- in minutes
	status VARCHAR(50) NOT NULL DEFAULT 'Pending',
		-- Pending, Confirmed, Preparing, Shipped, Delivered, Cancelled
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_orders_user_id ON orders (user_id);
