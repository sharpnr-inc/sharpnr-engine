package config

type TokenStructure struct {
	AccessToken  string
	RefreshToken string
}

type GenResponseBody[T any] struct {
	Error   *string `json:"error"`   // nil if no error
	Message string  `json:"message"` // description
	Data    *T      `json:"data"`    // nil if no data
}

// Success response
func NewSuccessResponse[T any](message string, data T) GenResponseBody[T] {
	return GenResponseBody[T]{
		Error:   nil,
		Message: message,
		Data:    &data,
	}
}

// Error response
func NewErrorResponse[T any](message string, err error) GenResponseBody[T] {
	errStr := err.Error()
	return GenResponseBody[T]{
		Error:   &errStr,
		Message: message,
		Data:    nil,
	}
}
