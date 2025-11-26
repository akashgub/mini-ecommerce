-- Admins Table for Admin Management
-- Admins create and manage products

CREATE TABLE IF NOT EXISTS admins (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    role VARCHAR(50) DEFAULT 'admin', -- admin, super_admin
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster lookups
CREATE INDEX IF NOT EXISTS idx_admins_username ON admins(username);
CREATE INDEX IF NOT EXISTS idx_admins_email ON admins(email);

-- Insert sample data
INSERT INTO admins (username, password, email, role) VALUES
('admin1', 'admin123', 'admin1@example.com', 'admin'),
('admin2', 'admin123', 'admin2@example.com', 'admin'),
('superadmin', 'admin123', 'superadmin@example.com', 'super_admin');

-- View all admins
SELECT * FROM admins;
