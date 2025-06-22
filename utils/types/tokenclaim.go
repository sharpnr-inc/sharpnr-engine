package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	EmployeeID     uint `json:"employee_id"`
	OrganizationID uint `json:"organization_id"`
	jwt.RegisteredClaims
}
