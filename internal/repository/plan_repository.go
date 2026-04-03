package repository

import (
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"gorm.io/gorm"
)

type PlanRepository struct {
	DB *gorm.DB
}

func NewPlanRepository(db *gorm.DB) *PlanRepository {
	return &PlanRepository{DB: db}
}

// Create a plan
func (r *PlanRepository) Create(plan *models.Plan) error {
	return r.DB.Create(plan).Error
}

// get all plan

func (r *PlanRepository) GetAll() ([]models.Plan, error) {
	var roles []models.Plan
	err := r.DB.Find(&roles).Error
	return roles, err
}

// get plan by  id
func (r *PlanRepository) GetByID(id uint) (*models.Plan, error) {
	var plan models.Plan
	err := r.DB.First(&plan, id).Error
	return &plan, err
}

// Update plan
func (r *PlanRepository) Update(plan *models.Plan) error {
	return r.DB.Save(plan).Error
}

// delete plan

func (r *PlanRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Role{}, id).Error
}

func (r *PlanRepository) GetRoleByTitle(title string) (*models.Role, error) {
	var plan models.Role
	err := r.DB.Where("role_title = ?", title).First(&plan).Error

	if err != nil {
		return nil, err
	}

	return &plan, nil
}
