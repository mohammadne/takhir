CREATE TABLE IF NOT EXISTS couriers(
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	phone_number VARCHAR(15) UNIQUE NOT NULL
	status VARCHAR(50) DEFAULT 'available', -- available, busy, inactive
	vehicle_info TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
