CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(128),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS credentials (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    method VARCHAR(50) CHECK (method IN ('phone', 'email', 'oauth')) NOT NULL,
    identifier VARCHAR(255) NOT NULL, -- Email address, phone number, or OAuth provider ID
    password_hash VARCHAR(255), -- Hashed password or null for OAuth
    oauth_provider VARCHAR(50), -- 'Google', 'Facebook' and etc only for OAuth
    verification_code VARCHAR(6), -- For phone/email only
    verified BOOLEAN DEFAULT FALSE, -- For phone/email to track verification status
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- we should have one credential per user 
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_credentials_user_id ON credentials (user_id);

-- Ensure unique combinations
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_credentials_method_identifier ON credentials (method, identifier);
