-- Products Table for Admin
-- Admin add products through this table

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    weight DECIMAL(10, 2) NOT NULL,
    colour VARCHAR(100),
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster queries
CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);

-- Insert sample data
INSERT INTO products (name, price, weight, colour, description) VALUES
('Apple', 150, 0.5, 'green', 'Fresh green apple'),
('Orange', 200, 0.4, 'orange', 'Fresh orange fruit'),
('Banana', 100, 0.3, 'yellow', 'Sweet yellow banana');

-- View all products
SELECT * FROM products;
