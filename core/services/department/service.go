package deptservices

import (
	"context"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/interfaces"
)

type departmentService struct {
	repo interfaces.DepartmentRepository
}

func DepartmentService(repo interfaces.DepartmentRepository) interfaces.DepartmentService {
	return &departmentService{repo: repo}
}

// ✅ Implement the method required by the interface
func (s *departmentService) AddDepartmentInOrg(ctx context.Context, req dto.AddDepartmentRequestDTO) error {
	// call repo, handle logic
	return nil
}

func (s *departmentService) RemoveDepartmentInOrg() {

}
