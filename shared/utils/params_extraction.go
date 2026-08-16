package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ExtractUUIDParam attempts to retrieve a parameter by key from path params first,
// falling back to query parameters, and parses it into a uuid.UUID.
func ExtractUUIDParam(ctx *gin.Context, key string) (uuid.UUID, bool) {
	val := ctx.Param(key)
	if val == "" {
		val = ctx.Query(key)
	}
	if val == "" {
		return uuid.Nil, false
	}

	id, err := uuid.Parse(val)
	if err != nil {
		return uuid.Nil, false
	}

	return id, true
}
