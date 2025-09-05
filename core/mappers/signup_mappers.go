package mappers

import (
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/core/models"
)

// FromSignUpREQToModel maps an Org_Signup_REQUEST_DTO to a TempOrgRegistrationData model.
// This function is typically used to convert the incoming signup request DTO
// into a format suitable for storage in the database or further processing.
//
// Parameters:
//   - req: The Org_Signup_REQUEST_DTO containing the organization's signup information.
//
// Returns:
//   - A pointer to TempOrgRegistrationData populated with data from the DTO.
//   - An error if mapping fails (currently always nil, but can be extended for validation).
func FromSignUpREQ_ToModel(req dto.Org_Signup_REQUEST_DTO) (*models.TempOrgRegistrationData, error) {
	return &models.TempOrgRegistrationData{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Mobile,
		OrgType:  req.ORG_TYPE,
		Password: req.Password,
	}, nil
}

// FromSignUp_RES_ToDTO maps TempOrgRegistrationData to Org_Signup_RESPONSE_DTO.
// This function converts the stored organization data into a response DTO
// and allows the service layer to pass in the generated authentication token
// and a custom message.
//
// Parameters:
//   - temp_org: Pointer to TempOrgRegistrationData containing the registered organization's data.
//   - token: Authentication token generated in the service layer.
//   - message: Custom message to include in the response (e.g., success message).
//
// Returns:
//   - Org_Signup_RESPONSE_DTO populated with the organization's email, token, and message.
//   - An error (currently always nil, but can be extended for validation or error handling).
func FromSignUp_RES_ToDTO(temp_org *models.TempOrgRegistrationData, token, message string) (dto.Org_Signup_RESPONSE_DTO, error) {
	return dto.Org_Signup_RESPONSE_DTO{
		Email:   temp_org.Email,
		Token:   token,
		Message: message,
	}, nil
}

// FromVerifyOtpRES_ToDTO maps an Organization model to Org_Verify_Otp_RESPONSE_DTO.
// This function is used after successfully verifying an OTP to construct
// a response DTO containing the organization's details and an authentication token.
//
// Parameters:
//   - org: Pointer to the Organization model containing the verified organization's data.
//   - token: Authentication token generated in the service layer.
//
// Returns:
//   - Org_Verify_Otp_RESPONSE_DTO populated with the organization's name, email,
//     unique Sharpnr ID, token, and organization type.
func FromVerifyOtpRES_ToDTO(org *models.Organization, token string) dto.Org_Verify_Otp_RESPONSE_DTO {
	return dto.Org_Verify_Otp_RESPONSE_DTO{
		Name:       org.Name,
		Email:      org.RootEmail,
		Sharpnr_Id: org.SharpnrOrgID,
		Token:      token,
		ORG_TYPE:   org.OrgOprationType,
	}
}
