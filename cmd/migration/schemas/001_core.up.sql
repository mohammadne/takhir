-- 
CREATE TABLE IF NOT EXISTS categories (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description TEXT
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_categories_unique_name ON categories (name);

CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	name VARCHAR(128),
	description TEXT,
    category_id INTEGER NOT NULL  REFERENCES categories(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_products_unique_name ON products (name);

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


---------------------------------------------------------- Insert sample Initial Data

INSERT INTO categories (name, description) 
VALUES 
    ('Laptops', 'Various types of laptops'),
	('Mobiles', 'Mobile phones'),
    ('Books', 'Printed books')
ON CONFLICT (name) DO NOTHING;

INSERT INTO products (name, description, category_id)
VALUES 
    -- Laptops
    ('MacBook Air M3', 'Apple MacBook Air with M3 chip, 256GB SSD, 16GB RAM', (SELECT id FROM categories WHERE name = 'Laptops')),
    ('MacBook Pro 16"', 'Apple MacBook Pro 16" with M1 Pro, 512GB SSD, 16GB RAM', (SELECT id FROM categories WHERE name = 'Laptops')),
    ('Dell XPS 13', 'Dell XPS 13, Intel i7, 16GB RAM, 512GB SSD', (SELECT id FROM categories WHERE name = 'Laptops')),

    -- Mobiles
    ('iPhone 16 Pro', 'Apple iPhone 16 Pro, 128GB, A18 Bionic chip', (SELECT id FROM categories WHERE name = 'Mobiles')),
    ('Samsung Galaxy S25 Ultra', 'Samsung Galaxy S25 Ultra, 256GB, Snapdragon 8 Gen 2', (SELECT id FROM categories WHERE name = 'Mobiles')),
    ('OnePlus 10 Pro', 'OnePlus 10 Pro, 128GB, Snapdragon 8 Gen 1', (SELECT id FROM categories WHERE name = 'Mobiles')),

    -- Books
    ('Science Fiction Book', 'Futuristic sci-fi novel', (SELECT id FROM categories WHERE name = 'Books')),
    ('History Book', 'A book about world history', (SELECT id FROM categories WHERE name = 'Books')),
    ('Programming Guide', 'Learn advanced programming concepts', (SELECT id FROM categories WHERE name = 'Books'))
ON CONFLICT (name) DO NOTHING;

INSERT INTO inventories (product_id, stock, price)
VALUES 
    -- Laptops
    ((SELECT id FROM products WHERE name = 'MacBook Air M3'), 10, 1000),
    ((SELECT id FROM products WHERE name = 'MacBook Pro 16"'), 8, 2500),
    ((SELECT id FROM products WHERE name = 'Dell XPS 13'), 12, 1500),

    -- Mobiles
    ((SELECT id FROM products WHERE name = 'iPhone 16 Pro'), 50, 1000),
    ((SELECT id FROM products WHERE name = 'Samsung Galaxy S25 Ultra'), 40, 1200),
    ((SELECT id FROM products WHERE name = 'OnePlus 10 Pro'), 35, 900),

    -- Books
    ((SELECT id FROM products WHERE name = 'Science Fiction Book'), 50, 25),
    ((SELECT id FROM products WHERE name = 'History Book'), 40, 30),
    ((SELECT id FROM products WHERE name = 'Programming Guide'), 35, 50)
ON CONFLICT (product_id) DO NOTHING;
