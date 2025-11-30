package admin

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"` // "admin", "super_admin"
}

// AdminResponse - safe response without password hash
type AdminResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type AdminLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// ToResponse converts Admin to AdminResponse (removes password)
func (a *Admin) ToResponse() *AdminResponse {
	return &AdminResponse{
		ID:       a.ID,
		Username: a.Username,
		Email:    a.Email,
		Role:     a.Role,
	}
}
