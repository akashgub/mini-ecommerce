package router

import (
	"gorm.io/gorm"

	"mini-ecommerce/internal/admin"
	"mini-ecommerce/internal/order"
	"mini-ecommerce/internal/product"
	"mini-ecommerce/internal/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize product repository, service, and handler
	productRepo := product.NewProductRepository(db)
	productService := product.NewProductService(productRepo)
	productHandler := product.NewProductHandler(productService)

	// Initialize admin repository, service, and handler
	adminRepo := admin.NewAdminRepository(db)
	adminService := admin.NewAdminService(adminRepo)
	adminHandler := admin.NewAdminHandler(adminService)

	// Initialize user repository, service, and handler
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// Initialize order repository, service, and handler
	orderRepo := order.NewOrderRepository(db)
	orderService := order.NewOrderService(orderRepo)
	orderHandler := order.NewOrderHandler(orderService, productRepo)

	// Product routes
	productRoutes := r.Group("/api/v1/products")
	{
		// Customer routes (public)
		productRoutes.GET("", productHandler.GetAllProducts)
		productRoutes.GET("/:id", productHandler.GetProductByID)

		// Admin routes (would need auth middleware in production)
		productRoutes.POST("", productHandler.CreateProduct)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

	// Admin routes
	adminRoutes := r.Group("/api/v1/admin")
	{
		adminRoutes.POST("/register", adminHandler.Register)
		adminRoutes.POST("/login", adminHandler.Login)
		adminRoutes.GET("", adminHandler.GetAllAdmins)
		adminRoutes.GET("/:id", adminHandler.GetAdminByID)
		adminRoutes.PUT("/:id", adminHandler.UpdateAdmin)
		adminRoutes.DELETE("/:id", adminHandler.DeleteAdmin)
	}

	// User routes
	userRoutes := r.Group("/api/v1/users")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
		userRoutes.GET("/profile/:id", userHandler.GetProfile)
		userRoutes.PUT("/profile/:id", userHandler.UpdateProfile)
		userRoutes.GET("", userHandler.GetAllUsers)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// Order routes
	orderRoutes := r.Group("/api/v1/orders")
	{
		// User can create and view their orders
		orderRoutes.POST("", orderHandler.CreateOrder)
		orderRoutes.GET("/user/:user_id", orderHandler.GetUserOrders)
		orderRoutes.GET("/:id", orderHandler.GetOrderByID)
		orderRoutes.DELETE("/:id", orderHandler.CancelOrder)

		// Admin only
		orderRoutes.GET("", orderHandler.GetAllOrders)
		orderRoutes.PUT("/:id/status", orderHandler.UpdateOrderStatus)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running"})
	})

	return r
}
