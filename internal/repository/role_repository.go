package repository

import (
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

// Create a role
func (r *RoleRepository) Create(role *models.Role) error {
	return r.DB.Create(role).Error
}

// get all role

func (r *RoleRepository) GetAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.Find(&roles).Error
	return roles, err
}

// get role by  id

func (r *RoleRepository) GetByID(id uint) (*models.Role, error) {
	var role models.Role
	err := r.DB.First(&role, id).Error
	return &role, err
}

// Update role
func (r *RoleRepository) Update(role *models.Role) error {
	return r.DB.Save(role).Error
}

// delete role

func (r *RoleRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Role{}, id).Error
}

func (r *RoleRepository) GetRoleByTitle(title string) (*models.Role, error) {
	var role models.Role
	err := r.DB.Where("role_title = ?", title).First(&role).Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}
