package hospitalitymodels

import "sharpnr-engine/core/models"

type Kitchen struct {
	ID             string
	OrganizationID string
	Organization   models.Organization
}
