-- 
CREATE TABLE IF NOT EXISTS translations (
	id SERIAL PRIMARY KEY,
	language VARCHAR(2) NOT NULL,
	translation TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
	id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS categories_translations (
	id SERIAL PRIMARY KEY,
	category_id INTEGER NOT NULL REFERENCES categories(id),
    translation_id INTEGER NOT NULL REFERENCES translations(id)
);

CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES categories(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products_translations (
	id SERIAL PRIMARY KEY,
    type VARCHAR(16) NOT NULL,
	product_id INTEGER NOT NULL REFERENCES products(id),
    translation_id INTEGER NOT NULL REFERENCES translations(id)
);

CREATE INDEX IF NOT EXISTS idx_products_category_id ON products (category_id);

CREATE TABLE IF NOT EXISTS inventories (
	id SERIAL PRIMARY KEY,
	product_id INTEGER NOT NULL REFERENCES products(id),
	stock INTEGER NOT NULL CHECK (stock >= 0),
	price INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_inventories_unique_product ON inventories (product_id);
