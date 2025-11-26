package order

import (
	"net/http"
	"strconv"

	"mini-ecommerce/internal/product"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service     OrderService
	productRepo product.ProductRepository
}

func NewOrderHandler(service OrderService, productRepo product.ProductRepository) *OrderHandler {
	return &OrderHandler{
		service:     service,
		productRepo: productRepo,
	}
}

// CreateOrder creates a new order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	order, err := h.service.CreateOrder(req, h.productRepo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetOrderByID retrieves an order by ID
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.service.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetUserOrders retrieves all orders for a user
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	orders, err := h.service.GetUserOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, []Order{})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetAllOrders retrieves all orders (admin only)
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, []Order{})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// UpdateOrderStatus updates order status (admin only)
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var req map[string]string
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	status, exists := req["status"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}

	order, err := h.service.UpdateOrderStatus(id, status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// CancelOrder cancels a pending order
func (h *OrderHandler) CancelOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	err = h.service.CancelOrder(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully"})
}
