package service

import (
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/repository"
)

type PermissionService struct {
	Repo *repository.PermissionsRepository
}

func NewPermissionService(repo *repository.PermissionsRepository) *PermissionService {
	return &PermissionService{Repo: repo}
}

func (s *PermissionService) CreatePlan(permission *models.Permission) error {
	return s.Repo.Create(permission)
}

func (s *PermissionService) GetAllRoles() ([]models.Permission, error) {
	return s.Repo.GetAll()
}

func (s *PermissionService) GetRoleByID(id uint) (*models.Permission, error) {
	return s.Repo.GetByID(id)
}

func (s *PermissionService) UpdateRole(permission *models.Permission) error {
	return s.Repo.Update(permission)
}

func (s *PermissionService) DeleteRole(id uint) error {
	return s.Repo.Delete(id)
}
