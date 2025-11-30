package admin

import (
	"errors"

	"mini-ecommerce/pkg/middleware"
)

type AdminService interface {
	Register(req AdminRegisterRequest) (*Admin, error)
	Login(req AdminLoginRequest) (map[string]interface{}, error)
	GetAdminByID(id int) (*Admin, error)
	UpdateAdmin(id int, admin *Admin) (*Admin, error)
	DeleteAdmin(id int) error
	GetAllAdmins() ([]Admin, error)
}

type adminService struct {
	repo AdminRepository
}

func NewAdminService(repo AdminRepository) AdminService {
	return &adminService{repo: repo}
}

func (s *adminService) Register(req AdminRegisterRequest) (*Admin, error) {
	// Check if admin already exists
	existing, _ := s.repo.FindByUsername(req.Username)
	if existing != nil {
		return nil, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := middleware.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	admin := &Admin{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "admin",
	}

	err = s.repo.Create(admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *adminService) Login(req AdminLoginRequest) (map[string]interface{}, error) {
	admin, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("admin not found")
	}

	// Verify password using bcrypt
	if !middleware.VerifyPassword(admin.Password, req.Password) {
		return nil, errors.New("invalid password")
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(admin.ID, admin.Email, admin.Username, admin.Role, "admin")
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return map[string]interface{}{
		"token": token,
		"admin": map[string]interface{}{
			"id":       admin.ID,
			"username": admin.Username,
			"email":    admin.Email,
			"role":     admin.Role,
		},
	}, nil
}

func (s *adminService) GetAdminByID(id int) (*Admin, error) {
	return s.repo.FindByID(id)
}

func (s *adminService) UpdateAdmin(id int, admin *Admin) (*Admin, error) {
	err := s.repo.Update(id, admin)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByID(id)
}

func (s *adminService) DeleteAdmin(id int) error {
	return s.repo.Delete(id)
}

func (s *adminService) GetAllAdmins() ([]Admin, error) {
	return s.repo.GetAll()
}
