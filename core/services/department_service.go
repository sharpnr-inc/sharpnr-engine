package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/sharpnr-inc/sharpnr-engine/core/constants"
	"github.com/sharpnr-inc/sharpnr-engine/core/interfaces"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
)

type departmentService struct {
	repo interfaces.DepartmentRepository
}

func NewDepartmentService(repo interfaces.DepartmentRepository) *departmentService {
	return &departmentService{repo: repo}
}

func (s *departmentService) CreateDefaultAccountingDepartment(orgUUID uuid.UUID) (*models.Department, error) {
	// In a real scenario, you might check if the department already exists for the organization

	dept := &models.Department{
		ID:             uuid.New(), // Corrected field name
		OrganizationID: orgUUID,
		Name:           string(constants.DepartmentTypeAccounting),
		Code:           "ACC", // Example code
		Type:           string(constants.DepartmentTypeAccounting), // Using constant for Type
		Description:    "Default Accounting Department",
		// CreatedAt and UpdatedAt will be set by GORM autoCreateTime/autoUpdateTime
	}

	if err := s.repo.Create(dept); err != nil {
		return nil, fmt.Errorf("failed to create accounting department in repo: %w", err)
	}

	fmt.Println("Default Accounting Department created.")
	return dept, nil
}
