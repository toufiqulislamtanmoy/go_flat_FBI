package service

import (
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Here you can add validation, e.g., email uniqueness
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
