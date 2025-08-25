/*
Package orgcontrollers provides HTTP controllers for handling organization-related
requests. It acts as the interface layer between incoming requests (e.g., API endpoints)
and the underlying business logic provided by services.

This package follows the dependency injection pattern, ensuring that controllers do not
directly depend on concrete implementations but instead rely on service interfaces.
This makes the code more testable, modular, and maintainable.
*/
package orgcontrollers

import "sharpnr-engine/core/interfaces"

// OrganisationController is responsible for handling requests related to
// organization registration and management. It delegates the core business
// logic to the OrgRegistrationService.
type OrganisationController struct {
	service interfaces.OrgRegistrationService
}

// RegisterOrgController creates a new OrganisationController instance.
//
// This function is part of the dependency injection setup. Instead of
// instantiating dependencies directly inside the controller, the service
// is injected from outside. This allows different implementations of
// OrgRegistrationService (e.g., mocks, stubs, real services) to be
// substituted during testing or runtime.
//
// Example usage:
//
//	service := services.NewOrgRegistrationService(repo)
//	controller := orgcontrollers.RegisterOrgController(service)
func RegisterOrgController(service interfaces.OrgRegistrationService) *OrganisationController {
	return &OrganisationController{service: service}
}
