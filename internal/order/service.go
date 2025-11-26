package order

import (
	"errors"

	"mini-ecommerce/internal/product"
)

type OrderService interface {
	CreateOrder(req CreateOrderRequest, productRepo product.ProductRepository) (*Order, error)
	GetOrderByID(id int) (*Order, error)
	GetUserOrders(userID int) ([]Order, error)
	GetAllOrders() ([]Order, error)
	UpdateOrderStatus(id int, status string) (*Order, error)
	CancelOrder(id int) error
}

type orderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(req CreateOrderRequest, productRepo product.ProductRepository) (*Order, error) {
	// Verify product exists
	prod, err := productRepo.FindByID(req.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	// Calculate total price
	totalPrice := prod.Price * float64(req.Quantity)

	order := &Order{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		TotalPrice: totalPrice,
		Status:    "pending",
	}

	err = s.repo.Create(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *orderService) GetOrderByID(id int) (*Order, error) {
	return s.repo.FindByID(id)
}

func (s *orderService) GetUserOrders(userID int) ([]Order, error) {
	return s.repo.FindByUserID(userID)
}

func (s *orderService) GetAllOrders() ([]Order, error) {
	return s.repo.FindAll()
}

func (s *orderService) UpdateOrderStatus(id int, status string) (*Order, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("order not found")
	}

	validStatuses := map[string]bool{"pending": true, "confirmed": true, "delivered": true}
	if !validStatuses[status] {
		return nil, errors.New("invalid status")
	}

	order.Status = status
	err = s.repo.Update(id, order)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(id)
}

func (s *orderService) CancelOrder(id int) error {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("order not found")
	}

	if order.Status != "pending" {
		return errors.New("only pending orders can be cancelled")
	}

	return s.repo.Delete(id)
}
