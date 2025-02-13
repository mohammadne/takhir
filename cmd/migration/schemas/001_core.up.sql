-- 
CREATE TABLE IF NOT EXISTS categories (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description TEXT
);

CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	name VARCHAR(128),
	description TEXT,
    category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE INDEX idx_products_category_id ON products (category_id);

CREATE TABLE IF NOT EXISTS inventories (
	id SERIAL PRIMARY KEY,
	product_id INTEGER NOT NULL, FOREIGN KEY (product_id) REFERENCES products(id),
	stock INTEGER NOT NULL CHECK (stock >= 0),
	price INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_inventories_unique_product ON inventories (product_id);
