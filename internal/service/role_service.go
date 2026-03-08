package service

import (
	"errors"

	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/repository"
)

type RoleService struct {
	Repo *repository.RoleRepository
}

func NewRoleService(repo *repository.RoleRepository) *RoleService {
	return &RoleService{Repo: repo}
}

func (s *RoleService) CreateRole(role *models.Role) error {
	allowedRoles := map[string]bool{
		"admin": true,
		"user":  true,
	}

	if !allowedRoles[role.RoleTitle] {
		return errors.New("invalid role title")
	}

	existing, err := s.Repo.GetRoleByTitle(role.RoleTitle)

	if err == nil && existing != nil {
		return errors.New("role already exists")
	}

	return s.Repo.Create(role)
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	return s.Repo.GetAll()
}

func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
	return s.Repo.GetByID(id)
}

func (s *RoleService) UpdateRole(role *models.Role) error {
	return s.Repo.Update(role)
}

func (s *RoleService) DeleteRole(id uint) error {
	return s.Repo.Delete(id)
}
