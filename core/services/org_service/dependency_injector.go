package orgservices

import "sharpnr-engine/core/interfaces"

// orgRegistrationService is a concrete implementation of the
// OrgRegistrationService interface. It acts as the service layer
// between the repository (data access) and the handlers/controllers.
//
// This struct embeds an IOrgRegistrationRepository dependency which
// provides the data persistence operations.
type orgRegistrationService struct {
	repo interfaces.IOrgRegistrationRepository
}

// OrgRegService is a constructor function (also known as a service injector).
// It accepts an implementation of IOrgRegistrationRepository and returns
// an OrgRegistrationService interface.
//
// This allows dependency injection, making the codebase more testable
// and decoupled from specific repository implementations.
//
// Example:
//
//	repo := repositories.NewOrgRegistrationRepository(db)
//	service := orgservices.OrgRegService(repo)
func OrgRegService(repo interfaces.IOrgRegistrationRepository) interfaces.OrgRegistrationService {
	return &orgRegistrationService{repo: repo}
}

type orgLoginService struct {
	repo interfaces.IOrgLoginRepository
}

func OrgLoginService(repo interfaces.IOrgLoginRepository) interfaces.OrgLoginService {
	return &orgLoginService{repo: repo}
}
