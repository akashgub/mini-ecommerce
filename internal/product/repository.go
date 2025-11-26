package product

import (
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *Product) error
	FindAll() ([]Product, error)
	FindByID(id int) (*Product, error)
	Update(id int, product *Product) error
	Delete(id int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id int) (*Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(id int, product *Product) error {
	return r.db.Model(&Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *productRepository) Delete(id int) error {
	return r.db.Delete(&Product{}, id).Error
}
