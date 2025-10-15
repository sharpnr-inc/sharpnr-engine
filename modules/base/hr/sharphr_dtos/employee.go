package sharphrdtos

type CheckIfEmployeeExist_REQ_DTO struct {
	EmployeeEmail string `json:"employee_email"`
}

type CheckIfEmployeeExist_RESPONSE_DTO struct {
	Exists string `json:"employee_exist"`
	Test   string `json:"query_email"`
	ORG    any    `json:"orguuid"`
	// Message string `json:"message"`
}
