# Database Setup Instructions

## Overview

এই প্রজেক্টে ৪টি মূল table আছে:

1. **admins** - Admin users যারা products manage করে
2. **products** - Products যা admin add করে
3. **users** - Customers যারা register এবং login করে
4. **orders** - Orders যা customers place করে

## Database Files

```
sql/
├── schema.sql                    # সম্পূর্ণ database schema
├── 01_create_products_table.sql  # Products table
├── 02_create_users_table.sql     # Users table
├── 03_create_orders_table.sql    # Orders table
└── 04_create_admins_table.sql    # Admins table
```

## Setup Methods

### Method 1: Using schema.sql (সহজ উপায়)

```bash
# PostgreSQL এ connect করুন
psql -U postgres -d ecommerce

# সম্পূর্ণ schema import করুন
\i sql/schema.sql

# সফল হলে এই output দেখবেন:
# table_name | count
# -----------+-------
# Admins     |     2
# Products   |     3
# Users      |     2
# Orders     |     2
```

### Method 2: Individual SQL Files

```bash
psql -U postgres -d ecommerce

# এক এক করে run করুন:
\i sql/01_create_products_table.sql
\i sql/02_create_users_table.sql
\i sql/03_create_orders_table.sql
\i sql/04_create_admins_table.sql
```

### Method 3: Command Line

```bash
# সম্পূর্ণ schema একসাথে apply করুন
psql -U postgres -d ecommerce -f sql/schema.sql
```

## Sample Data

### Admins (Product Management)
```
Username: admin1
Password: admin123
Email: admin1@example.com
Role: admin
```

### Users (Customers)
```
User 1:
Email: ahmed@example.com
Password: password123
Name: Ahmed Khan

User 2:
Email: fatima@example.com
Password: password123
Name: Fatima Begum
```

### Products
```
1. Apple - 150 taka, 0.5 kg, green
2. Orange - 200 taka, 0.4 kg, orange
3. Banana - 100 taka, 0.3 kg, yellow
```

## Table Relationships

```
admins
  └─ Creates and manages products

products
  ├─ Created by: admins
  └─ Ordered by: customers (via orders)

users (customers)
  ├─ Login with: email, password
  └─ Place: orders

orders
  ├─ user_id → references users(id)
  ├─ product_id → references products(id)
  └─ Managed by: admins
```

## Database Queries

### View All Admins
```sql
SELECT * FROM admins;
```

### View All Products
```sql
SELECT * FROM products;
```

### View All Users
```sql
SELECT * FROM users;
```

### View All Orders
```sql
SELECT * FROM orders;
```

### View Orders with Details
```sql
SELECT 
    o.id,
    u.name as customer_name,
    p.name as product_name,
    o.quantity,
    o.total_price,
    o.status,
    o.created_at
FROM orders o
JOIN users u ON o.user_id = u.id
JOIN products p ON o.product_id = p.id;
```

### View User's Orders
```sql
SELECT * FROM orders WHERE user_id = 1;
```

### View Orders by Status
```sql
SELECT * FROM orders WHERE status = 'pending';
SELECT * FROM orders WHERE status = 'confirmed';
SELECT * FROM orders WHERE status = 'delivered';
```

## API Testing with Sample Data

### Admin: Create Product
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mango",
    "price": 250,
    "weight": 0.6,
    "colour": "yellow",
    "description": "Sweet mango"
  }'
```

### User: Register
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Karim",
    "email": "karim@example.com",
    "phone": "01711223344",
    "password": "password123",
    "address": "Sylhet, Bangladesh"
  }'
```

### User: Login
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "ahmed@example.com",
    "password": "password123"
  }'
```

### User: Place Order
```bash
curl -X POST http://localhost:8080/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "product_id": 1,
    "quantity": 2
  }'
```

## Important Notes

- ✅ GORM automatically creates tables when you run `go run main.go`
- ✅ এই SQL files manual data insertion এর জন্য
- ✅ Foreign keys automatically maintain data integrity
- ✅ Indexes speed up queries
- ✅ Sample data দিয়ে testing করা সহজ হয়

## Troubleshooting

### Database already exists
```bash
# Drop and recreate database
psql -U postgres -c "DROP DATABASE IF EXISTS ecommerce;"
psql -U postgres -c "CREATE DATABASE ecommerce;"
```

### Permission denied
```bash
# Make sure PostgreSQL service is running
Get-Service postgresql-x64-18 # Windows
# Or on Mac/Linux:
sudo service postgresql status
```

### Cannot connect
```bash
# Check connection settings in .env
cat .env
```
