CREATE TABLE IF NOT EXISTS shippings (
	id SERIAL PRIMARY KEY,
	order_id INTEGER UNIQUE NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
	carrier VARCHAR(100) NOT NULL,  -- Name of the carrier
	tracking_code VARCHAR(100),  -- Tracking code provided by the carrier
	shipped_at TIMESTAMP,  -- The date when the order was shipped
	delivered_at TIMESTAMP,  -- The date when the order was delivered
	status VARCHAR(50), -- 'Pending', 'OnWay', 'Delievered', etc.
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_unique_shippings_tracking_code ON shippings (tracking_code);
