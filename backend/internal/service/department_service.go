package service

import (
    "errors"
    "internal-office-backend/internal/model"
    "internal-office-backend/internal/repository"
)

type DepartmentService interface {
    CreateDepartment(dept *model.Department) error
    GetAllDepartments() ([]model.Department, error)
    GetDepartmentByID(id uint) (*model.Department, error)
    UpdateDepartment(id uint, dept *model.Department) error
    DeleteDepartment(id uint) error
}

type departmentService struct {
    repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) DepartmentService {
    return &departmentService{repo}
}

func (s *departmentService) CreateDepartment(dept *model.Department) error {
    if dept.Name == "" {
        return errors.New("department name is required")
    }
    return s.repo.Create(dept)
}

func (s *departmentService) GetAllDepartments() ([]model.Department, error) {
    return s.repo.FindAll()
}

func (s *departmentService) GetDepartmentByID(id uint) (*model.Department, error) {
    return s.repo.FindByID(id)
}

func (s *departmentService) UpdateDepartment(id uint, dept *model.Department) error {
    existing, err := s.repo.FindByID(id)
    if err != nil {
        return err
    }

    existing.Name = dept.Name

    return s.repo.Update(existing)
}

func (s *departmentService) DeleteDepartment(id uint) error {
    return s.repo.Delete(id)
}