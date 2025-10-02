package dto

import "time"

type ErrorResponse struct {
	Message string `json:"message"`
	IsError bool   `json:"isError"`
}

type Response[T any] struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Error      ErrorResponse `json:"error"`
	Data       *T            `json:"data"` // <-- pointer here
	Timestamp  time.Time     `json:"timestamp"`
}
