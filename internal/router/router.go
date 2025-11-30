package router

import (
	"gorm.io/gorm"

	"mini-ecommerce/internal/admin"
	"mini-ecommerce/internal/order"
	"mini-ecommerce/internal/product"
	"mini-ecommerce/internal/user"
	"mini-ecommerce/pkg/middleware"

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

		// Admin routes (protected)
		adminProduct := productRoutes.Group("")
		adminProduct.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			adminProduct.POST("", productHandler.CreateProduct)
			adminProduct.PUT("/:id", productHandler.UpdateProduct)
			adminProduct.DELETE("/:id", productHandler.DeleteProduct)
		}
	}

	// Admin routes
	adminRoutes := r.Group("/api/v1/admin")
	{
		adminRoutes.POST("/register", adminHandler.Register)
		adminRoutes.POST("/login", adminHandler.Login)

		// Protected admin routes
		protectedAdmin := adminRoutes.Group("")
		protectedAdmin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			protectedAdmin.GET("", adminHandler.GetAllAdmins)
			protectedAdmin.GET("/:id", adminHandler.GetAdminByID)
			protectedAdmin.PUT("/:id", adminHandler.UpdateAdmin)
			protectedAdmin.DELETE("/:id", adminHandler.DeleteAdmin)
		}
	}

	// User routes
	userRoutes := r.Group("/api/v1/users")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)

		// Protected user routes
		protectedUser := userRoutes.Group("")
		protectedUser.Use(middleware.AuthMiddleware(), middleware.UserMiddleware())
		{
			protectedUser.GET("/profile/:id", userHandler.GetProfile)
			protectedUser.PUT("/profile/:id", userHandler.UpdateProfile)
		}

		// Admin only
		adminUser := userRoutes.Group("")
		adminUser.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			adminUser.GET("", userHandler.GetAllUsers)
			adminUser.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	// Order routes
	orderRoutes := r.Group("/api/v1/orders")
	{
		// Protected order routes
		protectedOrder := orderRoutes.Group("")
		protectedOrder.Use(middleware.AuthMiddleware())
		{
			protectedOrder.POST("", orderHandler.CreateOrder)
			protectedOrder.GET("/user/:user_id", orderHandler.GetUserOrders)
			protectedOrder.GET("/:id", orderHandler.GetOrderByID)
			protectedOrder.DELETE("/:id", orderHandler.CancelOrder)

			// Admin only
			adminOrder := orderRoutes.Group("")
			adminOrder.Use(middleware.AdminMiddleware())
			{
				adminOrder.GET("", orderHandler.GetAllOrders)
				adminOrder.PUT("/:id/status", orderHandler.UpdateOrderStatus)
			}
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running"})
	})

	return r
}
