package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register creates a new user
func (h *UserHandler) Register(c *gin.Context) {
	var req UserRegisterRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.service.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return only safe fields (no password hash)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":      user.ID,
			"name":    user.Name,
			"email":   user.Email,
			"phone":   user.Phone,
			"address": user.Address,
		},
	})
}

// Login authenticates a user
func (h *UserHandler) Login(c *gin.Context) {
	var req UserLoginRequest

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
		"user":    result["user"],
	})
}

// GetProfile retrieves user profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	profile := UserProfileResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, profile)
}

// UpdateProfile updates user profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req UserUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.service.UpdateUser(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user":    user.ToResponse(),
	})
}

// GetAllUsers retrieves all users (admin only)
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Convert to response format (without passwords)
	var responses []*UserResponse
	for _, user := range users {
		responses = append(responses, user.ToResponse())
	}

	c.JSON(http.StatusOK, responses)
}

// DeleteUser deletes a user (admin only)
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
