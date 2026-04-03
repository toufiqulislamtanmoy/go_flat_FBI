package service

import (
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/models"
	"github.com/toufiqulislamtanmoy/go_flat_FBI/internal/repository"
)

type PlanService struct {
	Repo *repository.PlanRepository
}

func NewPlanService(repo *repository.PlanRepository) *PlanService {
	return &PlanService{Repo: repo}
}

func (s *PlanService) CreatePlan(plan *models.Plan) error {
	return s.Repo.Create(plan)
}

func (s *PlanService) GetAllRoles() ([]models.Plan, error) {
	return s.Repo.GetAll()
}

func (s *PlanService) GetRoleByID(id uint) (*models.Plan, error) {
	return s.Repo.GetByID(id)
}

func (s *PlanService) UpdateRole(plan *models.Plan) error {
	return s.Repo.Update(plan)
}

func (s *PlanService) DeleteRole(id uint) error {
	return s.Repo.Delete(id)
}
