package departmentservice

import "sharpnr-engine/core/interfaces"

type departmentService struct {
	repo interfaces.DepartmentRepository
}

func DepartmentService(repo interfaces.DepartmentRepository) interfaces.DepartmentService {
	return &departmentService{repo: repo}
}
