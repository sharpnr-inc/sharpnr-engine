package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*
========================
Application Error Codes
========================
Stable identifiers for frontend & logs
*/
const (
	CodeValidationFailed = "VALIDATION_FAILED"
	CodeUnauthorized     = "UNAUTHORIZED"
	CodeForbidden        = "FORBIDDEN"
	CodeNotFound         = "NOT_FOUND"
	CodeConflict         = "CONFLICT"
	CodeInternalError    = "INTERNAL_ERROR"
)

/*
========================
API Error Model
========================
*/
type APIError struct {
	HTTPStatus int    `json:"status"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

/*
========================
Global API Response
========================
*/
type Response_[T any] struct {
	Success   bool      `json:"success"`
	Message   string    `json:"message,omitempty"`
	Data      T         `json:"data,omitempty"`
	Error     *APIError `json:"error,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Token     string    `json:"token,omitempty"`
}

/*
========================
Success Response
========================
*/
func Success[T any](
	ctx *gin.Context,
	status int,
	message string,
	data T,
) {
	ctx.JSON(status, Response_[T]{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	})
}

/*
========================
Failure Response
========================
*/
func Failure(
	ctx *gin.Context,
	status int,
	code string,
	err error,
) {
	msg := http.StatusText(status)
	if err != nil {
		msg = err.Error()
	}

	ctx.JSON(status, Response_[any]{
		Success: false,
		Error: &APIError{
			HTTPStatus: status,
			Code:       code,
			Message:    msg,
		},
		Timestamp: time.Now(),
	})
}

/*
========================
Validation Error
========================
*/
func ValidationError(ctx *gin.Context, err error) {
	fieldErrors := make(map[string]string)

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			fieldErrors[e.Field()] = e.Tag()
		}
	}

	ctx.JSON(http.StatusBadRequest, Response_[map[string]string]{
		Success: false,
		Error: &APIError{
			HTTPStatus: http.StatusBadRequest,
			Code:       CodeValidationFailed,
			Message:    "Invalid request payload",
		},
		Data:      fieldErrors,
		Timestamp: time.Now(),
	})
}
