CREATE TABLE IF NOT EXISTS inventories (
	id SERIAL PRIMARY KEY,
	item_id INTEGER NOT NULL, FOREIGN KEY (item_id) REFERENCES items(id),
	stock INTEGER NOT NULL CHECK (stock >= 0),
	price INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_inventories_unique_item ON inventories (item_id);
