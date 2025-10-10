package interfaces

import (
	dto "sharpnr-engine/core/dtos"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	LoginToAccount(ctx *gin.Context, req dto.LoginRequestDTO) (*dto.LoginResponseDTO, string, error)
	ResetPasswordOfUser()
}
