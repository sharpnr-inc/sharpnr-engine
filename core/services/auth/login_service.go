package services

import (
	"errors"
	dto "sharpnr-engine/core/dtos"
	"sharpnr-engine/shared/utils"

	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Login(req dto.OrgAuthRequestDTO) (*dto.OrgAuthResponseDTO, error) {
	org, err := s.repo.FindByMail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(org.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid Password")
	}

	token, err := utils.GenerateJWTToken(org.SharpnrOrgID, org.RootEmail)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &dto.OrgAuthResponseDTO{Token: token}, nil
}
