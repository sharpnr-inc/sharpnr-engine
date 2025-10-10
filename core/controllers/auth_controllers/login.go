package authcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := dto.Response[any]{
			StatusCode: http.StatusOK,
			Message:    "Some thing Went Wrong",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
			Token:     "",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	res, token, err := s.service.LoginToAccount(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response[any]{
			StatusCode: http.StatusUnauthorized,
			Message:    "Login failed",
			Error: dto.ErrorResponse{
				Message: err.Error(),
				IsError: true,
			},
			Data:      nil,
			Timestamp: time.Now(),
			Token:     "",
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response[dto.LoginResponseDTO]{
		StatusCode: 200,
		Message:    "Successfully Logged in",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      res,
		Timestamp: time.Now(),
		Token:     token,
	})

}
