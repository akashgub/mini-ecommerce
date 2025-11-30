package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service AdminService
}

func NewAdminHandler(service AdminService) *AdminHandler {
	return &AdminHandler{service: service}
}

// Register creates a new admin
func (h *AdminHandler) Register(c *gin.Context) {
	var req AdminRegisterRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	admin, err := h.service.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return only safe fields (no password hash)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin registered successfully",
		"admin": gin.H{
			"id":       admin.ID,
			"username": admin.Username,
			"email":    admin.Email,
			"role":     admin.Role,
		},
	})
}

// Login authenticates an admin
func (h *AdminHandler) Login(c *gin.Context) {
	var req AdminLoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   result["token"],
		"admin":   result["admin"],
	})
}

// GetAdminByID retrieves an admin by ID
func (h *AdminHandler) GetAdminByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	admin, err := h.service.GetAdminByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	c.JSON(http.StatusOK, admin.ToResponse())
}

// GetAllAdmins retrieves all admins
func (h *AdminHandler) GetAllAdmins(c *gin.Context) {
	admins, err := h.service.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch admins"})
		return
	}

	// Convert to response format (without passwords)
	var responses []*AdminResponse
	for _, admin := range admins {
		responses = append(responses, admin.ToResponse())
	}

	c.JSON(http.StatusOK, responses)
}

// UpdateAdmin updates an admin
func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	var admin Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedAdmin, err := h.service.UpdateAdmin(id, &admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin updated successfully",
		"admin":   updatedAdmin.ToResponse(),
	})
}

// DeleteAdmin deletes an admin
func (h *AdminHandler) DeleteAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	err = h.service.DeleteAdmin(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}
