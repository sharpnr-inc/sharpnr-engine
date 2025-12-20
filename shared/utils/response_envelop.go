package utils

import (
	dto "sharpnr-engine/core/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

func SuccessResponse[T any](ctx *gin.Context, status int, message string, data T) {
	ctx.JSON(status, dto.Response[T]{
		StatusCode: status,
		Message:    message,
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &data,
		Timestamp: time.Now(),
	})
}

func ErrorResponse(ctx *gin.Context, status int, message string, err error) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	ctx.JSON(status, dto.Response[any]{
		StatusCode: status,
		Message:    message,
		Error: dto.ErrorResponse{
			Message: errMsg,
			IsError: true,
		},
		Data:      nil,
		Timestamp: time.Now(),
	})
}
