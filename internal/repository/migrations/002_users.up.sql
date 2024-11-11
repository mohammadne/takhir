CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	name VARCHAR(30),
	username VARCHAR(30),
	email VARCHAR(60) NOT NULL,
	password VARCHAR(30) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX users_username_idx ON users (username);

CREATE INDEX users_email_idx ON users (email);
