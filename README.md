# Mini E-Commerce Platform

A simple e-commerce API built with Go, Gin, and PostgreSQL where customers can browse products and admins can manage inventory.

## Features

- **Admin Product Management**: Create, read, update, and delete products
- **Customer Shopping**: Browse all available products with detailed information
- Product Information: ID, Name, Price, Weight (kg), Description
- RESTful API with Gin web framework
- PostgreSQL database with GORM ORM

## Project Structure

```
.
├── main.go                      # Application entry point
├── config/
│   └── config.go                # Configuration management
├── db/
│   └── database.go              # Database connection & migration
├── internal/
│   ├── product/
│   │   ├── handler.go           # HTTP handlers for products
│   │   ├── model.go             # Product data model
│   │   ├── repository.go        # Database operations
│   │   └── service.go           # Business logic
│   └── router/
│       └── router.go            # Route definitions
├── go.mod
└── go.sum
```

## Prerequisites

- Go 1.24.5 or higher
- PostgreSQL 12 or higher
- Git

## Setup

### 1. Clone the Repository

```bash
git clone <your-repo-url>
cd mini-ecommerce
```

### 2. Configure Environment

Copy `.env.example` to `.env` and update the database credentials:

```bash
cp .env.example .env
```

Update `.env` with your PostgreSQL connection details:

```env
PORT=8080
DB_DRIVER=postgres
DB_USER=postgres
DB_PASS=your_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=ecommerce
```

### 3. Create Database

```bash
createdb ecommerce
```

### 4. Install Dependencies

```bash
go mod download
```

### 5. Run the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Health Check
- `GET /health` - Check if server is running

### Products (Customer - Public)
- `GET /api/v1/products` - Get all products
- `GET /api/v1/products/:id` - Get a specific product

### Products (Admin - Management)
- `POST /api/v1/products` - Create a new product
- `PUT /api/v1/products/:id` - Update a product
- `DELETE /api/v1/products/:id` - Delete a product

## Example Requests

### Get All Products
```bash
curl http://localhost:8080/api/v1/products
```

### Get Product by ID
```bash
curl http://localhost:8080/api/v1/products/1
```

### Create a Product (Admin)
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Apple",
    "price": 150,
    "weight": 0.5,
    "description": "Apple is a green color fruit"
  }'
```

### Update a Product (Admin)
```bash
curl -X PUT http://localhost:8080/api/v1/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Red Apple",
    "price": 180,
    "weight": 0.6,
    "description": "Apple is a red color fruit"
  }'
```

### Delete a Product (Admin)
```bash
curl -X DELETE http://localhost:8080/api/v1/products/1
```

## Product Model

```json
{
  "id": 1,
  "name": "Apple",
  "price": 150,
  "weight": 0.5,
  "description": "Apple is a green color fruit",
  "created_at": "2025-11-25T10:30:00Z",
  "updated_at": "2025-11-25T10:30:00Z"
}
```

## Database Schema

### Products Table
```sql
CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  weight DECIMAL(10, 2) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Future Enhancements

- User authentication and authorization
- Shopping cart functionality
- Order management system
- Payment integration
- Product categories
- User reviews and ratings
- Inventory tracking
- Admin dashboard

## Technology Stack

- **Language**: Go 1.24.5
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Database Driver**: pgx (PostgreSQL)

## License

MIT License

## Author

Mini E-Commerce Development Team
