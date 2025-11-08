package sharphrservice

import sharphrdtos "sharpnr-engine/modules/hr/sharphr_dtos"

// AddEmployeeInOrg implements sharphrinterface.SharpHrEmployeeService.
func (r *sharpHrService) AddEmployeeInOrg() error {
	panic("unimplemented")
}

func (r *sharpHrService) GetEmployeeDetailById() error {
	panic("unimplemented")
}

func (r *sharpHrService) ListAllEmployeeInOrg() error {
	panic("unimplemented")
}

func (r *sharpHrService) GetEmployeeDetailByEmail(orgId, query_email string, org_uuid any) (*sharphrdtos.CheckIfEmployeeExist_RESPONSE_DTO, error) {
	return &sharphrdtos.CheckIfEmployeeExist_RESPONSE_DTO{
		Exists: orgId,
		Test:   query_email,
		ORG:    org_uuid,
	}, nil
}
