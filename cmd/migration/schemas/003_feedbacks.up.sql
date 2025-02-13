CREATE TABLE IF NOT EXISTS feedback (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(10) NOT NULL CHECK (type IN ('comment', 'rating')),
    content TEXT, -- Comment text or null for ratings
    rating INTEGER CHECK (rating BETWEEN 1 AND 5), -- Rating value or null for comments
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_feedback_product_id ON feedback (product_id);
