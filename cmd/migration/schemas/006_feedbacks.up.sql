CREATE TABLE IF NOT EXISTS feedback (
    id SERIAL PRIMARY KEY,
    item_id INTEGER NOT NULL, FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL, FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(10) NOT NULL CHECK (type IN ('comment', 'rating')),
    content TEXT, -- Comment text or null for ratings
    rating INTEGER CHECK (rating BETWEEN 1 AND 5), -- Rating value or null for comments
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_feedback_item_id ON feedback (item_id);
