package user

import (
	"errors"
)

type UserService interface {
	Register(req UserRegisterRequest) (*User, error)
	Login(req UserLoginRequest) (*User, error)
	GetUserByID(id int) (*User, error)
	UpdateUser(id int, req UserUpdateRequest) (*User, error)
	DeleteUser(id int) error
	GetAllUsers() ([]User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(req UserRegisterRequest) (*User, error) {
	// Check if user already exists
	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password, // In production, hash the password
		Address:  req.Address,
	}

	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Login(req UserLoginRequest) (*User, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// In production, use bcrypt to compare passwords
	if user.Password != req.Password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s *userService) GetUserByID(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id int, req UserUpdateRequest) (*User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Address != "" {
		user.Address = req.Address
	}

	err = s.repo.Update(id, user)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(id)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAll()
}
