-- Users Table for Customers
-- Users register and login through this table

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL,
    password VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster lookups
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Insert sample data
INSERT INTO users (name, email, phone, password, address) VALUES
('Ahmed Khan', 'ahmed@example.com', '01712345678', 'password123', 'Dhaka, Bangladesh'),
('Fatima Begum', 'fatima@example.com', '01987654321', 'password123', 'Chittagong, Bangladesh'),
('Rajib Kumar', 'rajib@example.com', '01556789012', 'password123', 'Sylhet, Bangladesh');

-- View all users
SELECT * FROM users;
