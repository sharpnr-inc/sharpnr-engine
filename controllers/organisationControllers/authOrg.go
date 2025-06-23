package controllers

import "github.com/gin-gonic/gin"

// RegisterOrganization handles registration of a new organization.
//
// This endpoint should be excluded from authentication middleware.
//
// Expected JSON Payload:
//   {
//     "name": "Acme Corp",
//     "email": "admin@acme.com",
//     "password": "secret",
//     ...
//   }
//
// Returns:
//   - 201 Created on success
//   - 400 Bad Request on validation failure
//   - 500 Internal Server Error on failure
func RegisterOrganization(c *gin.Context) {
	// Implementation here
}

// OrganizationLogin handles login for organizations.
//
// Accepts either email or username and a password.
// This endpoint should be excluded from authentication middleware.
//
// Expected JSON Payload:
//   {
//     "email": "admin@acme.com",
//     "password": "secret"
//   }
//
// Returns:
//   - 200 OK with JWT token
//   - 401 Unauthorized if credentials are invalid
func OrganizationRootLogin(c *gin.Context) {
	// Implementation here
}
