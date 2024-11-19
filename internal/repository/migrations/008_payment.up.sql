CREATE TABLE IF NOT EXISTS payments (
	id SERIAL PRIMARY KEY,
	order_id INTEGER NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id),
	method VARCHAR(50), -- 'Credit Card', 'PayPal', etc.
	status VARCHAR(50), -- 'Pending', 'Completed', 'Failed', etc.
	amount INTEGER NOT NULL,
	date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_unique_payments_id_order_id ON payments (id, order_id);
