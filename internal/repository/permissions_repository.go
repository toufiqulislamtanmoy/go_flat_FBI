package repository

import (
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"gorm.io/gorm"
)

type PermissionsRepository struct {
	DB *gorm.DB
}

func NewPermissionsRepository(db *gorm.DB) *PermissionsRepository {
	return &PermissionsRepository{DB: db}
}

// Create a permission
func (r *PermissionsRepository) Create(permission *models.Permission) error {
	return r.DB.Create(permission).Error
}

// get all permission

func (r *PermissionsRepository) GetAll() ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.DB.Find(&permissions).Error
	return permissions, err
}

// get permission by  id
func (r *PermissionsRepository) GetByID(id uint) (*models.Permission, error) {
	var permission models.Permission
	err := r.DB.First(&permission, id).Error
	return &permission, err
}

// Update permission
func (r *PermissionsRepository) Update(permission *models.Permission) error {
	return r.DB.Save(permission).Error
}

// delete permission

func (r *PermissionsRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Permission{}, id).Error
}

func (r *PermissionsRepository) GetRoleByTitle(title string) (*models.Permission, error) {
	var permission models.Permission
	err := r.DB.Where("role_title = ?", title).First(&permission).Error

	if err != nil {
		return nil, err
	}

	return &permission, nil
}
