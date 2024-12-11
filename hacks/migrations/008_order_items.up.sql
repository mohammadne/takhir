CREATE TABLE IF NOT EXISTS order_items (
	id SERIAL PRIMARY KEY,
	order_id INTEGER NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id),
	item_id INTEGER NOT NULL, FOREIGN KEY (item_id) REFERENCES items(id),
	price INTEGER NOT NULL, -- Price of the item at the time of checkout
	quantity INTEGER NOT NULL CHECK (quantity > 0),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX idx_unique_order_items_order_id_item_id ON order_items (order_id, item_id);