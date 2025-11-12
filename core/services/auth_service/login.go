package authservice

import (
	"fmt"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/pkg/jwtpack"
	"sharpnr-engine/shared/utils"

	"github.com/gin-gonic/gin"
)

func (r *loginService) LoginToAccount(ctx *gin.Context, req dto.LoginRequestDTO) (*dto.LoginResponseDTO, string, error) {
	if req.Email == "" || req.Password == "" {
		return nil, "", fmt.Errorf("email and password are required")
	}
	usr, found, err := r.hrEmpRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", fmt.Errorf("failed to find employee: %w", err)
	}
	if !found {
		return nil, "", fmt.Errorf("employee not found")
	}

	// Validate Password

	usrOrg, err := r.orgRepo.GetOrg(ctx, usr.OrganizationID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to fetch organization: %w", err)
	}
	if usrOrg == nil {
		return nil, "", fmt.Errorf("organization not found")
	}

	if !utils.CheckPasswordHash(req.Password, usrOrg.Password) {
		return nil, "", fmt.Errorf("invalid credentials")
	}

	resp := dto.LoginResponseDTO{
		OrgId:         usrOrg.SharpnrID,
		OrgType:       string(usrOrg.OrgOperationType),
		EmployeeID:    usr.ID,
		EmployeeEmail: usr.Email,
	}
	empIdStr := resp.EmployeeID.String()
	newToken, err := jwtpack.GenerateJwtToken(resp.EmployeeEmail, jwtpack.TokenTypeAuth, resp.OrgId, resp.OrgType, &empIdStr)
	if err != nil {
		return nil, "", err
	}

	return &resp, newToken, nil
}
