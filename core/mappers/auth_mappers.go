package mappers

import (
	dto "sharpnr-engine/core/dtos"
	sharphr_models "sharpnr-engine/modules/hr/hrModels"
)

// This Function Takes the the Login details email, password and return a model
// Since we have common login for org and employee it returns one of two models
// dto.LoginResponseDTO_Org or dto.LoginResponseDTO_Employee
func FromLoginRequestToModel(req dto.LoginRequestDTO) *sharphr_models.Employee {
	return &sharphr_models.Employee{}

}

func FromRequestDtoToModel(user *sharphr_models.Employee) dto.LoginResponseDTO {
	return dto.LoginResponseDTO{}
}
