CREATE TABLE IF NOT EXISTS orders (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	status VARCHAR(32) NOT NULL, -- Pending, Confirmed, Preparing, Shipped, Delivered, Cancelled
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders (user_id);

CREATE TABLE IF NOT EXISTS order_products (
	id SERIAL PRIMARY KEY,
	order_id INTEGER NOT NULL REFERENCES orders(id),
	product_id INTEGER NOT NULL REFERENCES products(id),
	price INTEGER NOT NULL, -- Price of the product at the time of checkout
	quantity INTEGER NOT NULL CHECK (quantity > 0),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_order_products_order_id_product_id ON order_products (order_id, product_id);
