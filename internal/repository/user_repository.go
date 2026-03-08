package repository

import (
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create a new user
func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

// Get all users
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Preload("Role").Find(&users).Error
	return users, err
}

// Get a user by ID
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.Preload("Role").First(&user, id).Error
	return &user, err
}

// Update user
func (r *UserRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

// Delete user
func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}

// user login
func (r *UserRepository) Login(user *models.User) (*models.User, error) {
	var singleUser models.User
	err := r.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&singleUser).Error
	return &singleUser, err
}
