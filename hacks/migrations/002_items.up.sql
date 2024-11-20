CREATE TABLE IF NOT EXISTS items (
	id SERIAL PRIMARY KEY,
	name VARCHAR(128),
	description TEXT,
    category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);
