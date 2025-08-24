package departmentservice

import (
	"context"
	dto "sharpnr-engine/core/dtos"
)

// ✅ Implement the method required by the interface
func (s *departmentService) AddDepartmentInOrg(ctx context.Context, req dto.AddDepartmentRequestDTO) error {
	// call repo, handle logic
	return nil
}

func (s *departmentService) RemoveDepartmentInOrg() {

}
