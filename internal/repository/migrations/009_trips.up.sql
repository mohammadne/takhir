CREATE TABLE IF NOT EXISTS trips (
	id SERIAL PRIMARY KEY,
	order_id INTEGER UNIQUE NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
	status VARCHAR(50) CHECK (status IN ('ASSIGNED', 'PICKED', 'DELIVERED'))
		NOT NULL DEFAULT 'ASSIGNED',
	assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	picked_at TIMESTAMP,
	delivered_at TIMESTAMP
);
