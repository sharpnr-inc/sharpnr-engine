package mappers

import (
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
)

// This Function Takes the the Login details email, password and return a model
// Since we have common login for org and employee it returns one of two models
// dto.LoginResponseDTO_Org or dto.LoginResponseDTO_Employee
func FromLoginRequestToModel(req dto.LoginRequestDTO) *models.Employee {
	return &models.Employee{}

}

func FromRequestDtoToModel(user *models.Employee) dto.LoginResponseDTO {
	return dto.LoginResponseDTO{}
}
