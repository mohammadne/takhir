CREATE TABLE IF NOT EXISTS delay_reports (
    id SERIAL PRIMARY KEY,
    order_id INTEGER UNIQUE NOT NULL, FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    delay_reason TEXT,
    estimated_delivery_time TIMESTAMP,
    processed BOOLEAN DEFAULT FALSE,
    report_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
