package order

import "time"

type Order struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	TotalPrice float64  `json:"total_price"`
	Status    string    `json:"status"` // pending, confirmed, delivered
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    float64 `json:"price"`
	Quantity int    `json:"quantity"`
	Total    float64 `json:"total"`
}

type CreateOrderRequest struct {
	UserID    int `json:"user_id" binding:"required,gt=0"`
	ProductID int `json:"product_id" binding:"required,gt=0"`
	Quantity  int `json:"quantity" binding:"required,gt=0"`
}

type OrderResponse struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	ProductID  int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
