CREATE TABLE IF NOT EXISTS order_items(
	id SERIAL PRIMARY KEY,
	order_id INTEGER NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id),
	item_id INTEGER NOT NULL, FOREIGN KEY (item_id) REFERENCES items(id),
	quantity INTEGER NOT NULL CHECK (quantity > 0)
);

CREATE UNIQUE INDEX idx_unique_order_items_order_id_item_id ON order_items (order_id, item_id);
