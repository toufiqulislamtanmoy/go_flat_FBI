package service

import (
	"errors"
	"fmt"

	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/repository"
)

type UserService struct {
	Repo     *repository.UserRepository
	RoleRepo *repository.RoleRepository
}

func NewUserService(repo *repository.UserRepository, roleRepo *repository.RoleRepository) *UserService {
	return &UserService{Repo: repo, RoleRepo: roleRepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	fmt.Println("user-->",user)
	if user.RoleID == 0 {
		return errors.New("role_id is required")
	}

	_, err := s.RoleRepo.GetByID(user.RoleID)
	if err != nil {
		return fmt.Errorf("role_id invalid: %w", err)
	}

	// Here you can add additional validation, e.g., email uniqueness
	return s.Repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.Delete(id)
}

func (s *UserService) Login(user *models.User) (*models.User, error) {
	return s.Repo.Login(user)
}
