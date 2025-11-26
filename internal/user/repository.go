package user

import "gorm.io/gorm"

type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id int) (*User, error)
	Update(id int, user *User) error
	Delete(id int) error
	GetAll() ([]User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id int) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(id int, user *User) error {
	return r.db.Model(&User{}).Where("id = ?", id).Updates(user).Error
}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *userRepository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}
