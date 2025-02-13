CREATE TABLE IF NOT EXISTS inventories (
	id SERIAL PRIMARY KEY,
	product_id INTEGER NOT NULL, FOREIGN KEY (product_id) REFERENCES products(id),
	stock INTEGER NOT NULL CHECK (stock >= 0),
	price INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_inventories_unique_product ON inventories (product_id);
