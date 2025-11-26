package admin

import (
	"errors"
)

type AdminService interface {
	Register(req AdminRegisterRequest) (*Admin, error)
	Login(req AdminLoginRequest) (*Admin, error)
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

	admin := &Admin{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // In production, hash the password
		Role:     "admin",
	}

	err := s.repo.Create(admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *adminService) Login(req AdminLoginRequest) (*Admin, error) {
	admin, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("admin not found")
	}

	// In production, use bcrypt to compare passwords
	if admin.Password != req.Password {
		return nil, errors.New("invalid password")
	}

	return admin, nil
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
