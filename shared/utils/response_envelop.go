package utils

import (
	dto "sharpnr-engine/core/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

type Error_response struct {
	Message string `json:"message"`
	IsError bool   `json:"isError"`
}

type Response[T any] struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Error      Error_response `json:"error"`
	Data       *T             `json:"data"` // <-- pointer here
	Timestamp  time.Time      `json:"timestamp"`
	Token      string         `json:"token,omitempty"` // ✅ Optional field
}

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
