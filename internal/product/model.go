package product

import "time"

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"index"`
	Price       float64   `json:"price"`
	Weight      float64   `json:"weight"` // in kg
	Colour      string    `json:"colour"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Weight      float64 `json:"weight" binding:"required,gt=0"`
	Colour      string  `json:"colour" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Weight      float64 `json:"weight"`
	Colour      string  `json:"colour"`
	Description string  `json:"description"`
}
