package repository

import (
	"internal-office-backend/internal/model"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
    Create(dept *model.Department) error
    FindAll() ([]model.Department, error)
    FindByID(id uint) (*model.Department, error)
    Update(dept *model.Department) error
    Delete(id uint) error
}

type departmentRepository struct {
    db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
    return &departmentRepository{db}
}

func (r *departmentRepository) Create(dept *model.Department) error {
    return r.db.Create(dept).Error
}

func (r *departmentRepository) FindAll() ([]model.Department, error) {
    var departments []model.Department
    err := r.db.Find(&departments).Error
    return departments, err
}

func (r *departmentRepository) FindByID(id uint) (*model.Department, error) {
    var department model.Department
    err := r.db.First(&department, id).Error
    if err != nil {
        return nil, err
    }
    return &department, nil
}

func (r *departmentRepository) Update(dept *model.Department) error {
    return r.db.Save(dept).Error
}

func (r *departmentRepository) Delete(id uint) error {
    return r.db.Delete(&model.Department{}, id).Error
}