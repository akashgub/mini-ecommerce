-- Orders Table for Order Management
-- Tracks all customer orders

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

-- Create indexes for faster queries
CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_product_id ON orders(product_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);

-- Insert sample data
INSERT INTO orders (user_id, product_id, quantity, total_price, status) VALUES
(1, 1, 2, 300, 'pending'),
(2, 2, 3, 600, 'confirmed'),
(3, 3, 1, 100, 'delivered');

-- View all orders
SELECT * FROM orders;

-- View user's orders
SELECT o.id, o.user_id, u.name, u.email, o.product_id, p.name as product_name, 
       o.quantity, o.total_price, o.status, o.created_at
FROM orders o
JOIN users u ON o.user_id = u.id
JOIN products p ON o.product_id = p.id;
