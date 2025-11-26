package order

import "gorm.io/gorm"

type OrderRepository interface {
	Create(order *Order) error
	FindByID(id int) (*Order, error)
	FindByUserID(userID int) ([]Order, error)
	FindAll() ([]Order, error)
	Update(id int, order *Order) error
	Delete(id int) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) FindByID(id int) (*Order, error) {
	var order Order
	err := r.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) FindByUserID(userID int) ([]Order, error) {
	var orders []Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) FindAll() ([]Order, error) {
	var orders []Order
	err := r.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) Update(id int, order *Order) error {
	return r.db.Model(&Order{}).Where("id = ?", id).Updates(order).Error
}

func (r *orderRepository) Delete(id int) error {
	return r.db.Delete(&Order{}, id).Error
}
