CREATE TABLE IF NOT EXISTS shippings (
	id SERIAL PRIMARY KEY,
	order_id INTEGER UNIQUE NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
	carrier VARCHAR(100) NOT NULL,  -- Name of the carrier
	tracking_code VARCHAR(100),  -- Tracking code provided by the carrier
	shipped_at TIMESTAMP,  -- The date when the order was shipped
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
