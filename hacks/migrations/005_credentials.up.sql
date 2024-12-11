CREATE TABLE IF NOT EXISTS credentials (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL, FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    method VARCHAR(50) CHECK (method IN ('phone', 'email', 'oauth')) NOT NULL,
    identifier VARCHAR(255) NOT NULL, -- Email address, phone number, or OAuth provider ID
    password_hash VARCHAR(255), -- Hashed password or null for OAuth
    oauth_provider VARCHAR(50), -- 'Google', 'Facebook' and etc only for OAuth
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- we should have one credential per user 
CREATE UNIQUE INDEX idx_unique_credentials_user_id ON credentials (user_id);

-- Ensure unique combinations
CREATE UNIQUE INDEX idx_unique_credentials_method_identifier ON credentials (method, identifier);
