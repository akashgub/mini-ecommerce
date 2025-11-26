-- Mini E-Commerce Database Schema
-- Complete database setup for the project

-- ============================================
-- 1. ADMINS TABLE
-- ============================================
CREATE TABLE IF NOT EXISTS admins (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    role VARCHAR(50) DEFAULT 'admin', -- admin, super_admin
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_admins_username ON admins(username);
CREATE INDEX IF NOT EXISTS idx_admins_email ON admins(email);

-- ============================================
-- 2. PRODUCTS TABLE
-- ============================================
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

CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);

-- ============================================
-- 3. USERS TABLE
-- ============================================
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

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- ============================================
-- 4. ORDERS TABLE
-- ============================================
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    total_price DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending', -- pending, confirmed, delivered
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_product_id ON orders(product_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);

-- ============================================
-- SAMPLE DATA
-- ============================================

-- Insert Admins
INSERT INTO admins (username, password, email, role) VALUES
('admin1', 'admin123', 'admin1@example.com', 'admin'),
('superadmin', 'admin123', 'superadmin@example.com', 'super_admin');

-- Insert Products
INSERT INTO products (name, price, weight, colour, description) VALUES
('Apple', 150, 0.5, 'green', 'Fresh green apple'),
('Orange', 200, 0.4, 'orange', 'Fresh orange fruit'),
('Banana', 100, 0.3, 'yellow', 'Sweet yellow banana');

-- Insert Users
INSERT INTO users (name, email, phone, password, address) VALUES
('Ahmed Khan', 'ahmed@example.com', '01712345678', 'password123', 'Dhaka, Bangladesh'),
('Fatima Begum', 'fatima@example.com', '01987654321', 'password123', 'Chittagong, Bangladesh');

-- Insert Orders
INSERT INTO orders (user_id, product_id, quantity, total_price, status) VALUES
(1, 1, 2, 300, 'pending'),
(2, 2, 3, 600, 'confirmed');

-- ============================================
-- VERIFY DATA
-- ============================================
SELECT 'Admins' as table_name, COUNT(*) as count FROM admins
UNION ALL
SELECT 'Products', COUNT(*) FROM products
UNION ALL
SELECT 'Users', COUNT(*) FROM users
UNION ALL
SELECT 'Orders', COUNT(*) FROM orders;
