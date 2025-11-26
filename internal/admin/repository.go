package admin

import "gorm.io/gorm"

type AdminRepository interface {
	Create(admin *Admin) error
	FindByUsername(username string) (*Admin, error)
	FindByID(id int) (*Admin, error)
	Update(id int, admin *Admin) error
	Delete(id int) error
	GetAll() ([]Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) Create(admin *Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) FindByUsername(username string) (*Admin, error) {
	var admin Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) FindByID(id int) (*Admin, error) {
	var admin Admin
	err := r.db.First(&admin, id).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) Update(id int, admin *Admin) error {
	return r.db.Model(&Admin{}).Where("id = ?", id).Updates(admin).Error
}

func (r *adminRepository) Delete(id int) error {
	return r.db.Delete(&Admin{}, id).Error
}

func (r *adminRepository) GetAll() ([]Admin, error) {
	var admins []Admin
	err := r.db.Find(&admins).Error
	return admins, err
}
