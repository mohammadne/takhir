CREATE TABLE IF NOT EXISTS payments (
	id SERIAL PRIMARY KEY,
	order_id INTEGER NOT NULL REFERENCES orders(id),
	method VARCHAR(50), -- 'Credit Card', 'PayPal', etc.
	status VARCHAR(50), -- 'Pending', 'Completed', 'Failed', etc.
	amount INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_payments_id_order_id ON payments (id, order_id);
