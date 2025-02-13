CREATE TABLE IF NOT EXISTS order_products (
	id SERIAL PRIMARY KEY,
	order_id INTEGER NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id),
	product_id INTEGER NOT NULL, FOREIGN KEY (product_id) REFERENCES products(id),
	price INTEGER NOT NULL, -- Price of the product at the time of checkout
	quantity INTEGER NOT NULL CHECK (quantity > 0),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX idx_unique_order_products_order_id_product_id ON order_products (order_id, product_id);
