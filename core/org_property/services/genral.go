package org_property_services

import (
	"context"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	org_property_models "sharpnr-engine/core/org_property/models"
	"time"

	"github.com/google/uuid"
)

func (c *organizationPropertyService) CreatePropertyDefault(ctx context.Context, input org_property_contracts.CreateOrgPropertyInput) error {

	property := &org_property_models.OrganizationProperty{
		ID:             uuid.New(),
		OrganizationID: input.OrgID,
		Code:           "p1",
		Name:           input.Name,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := c.orgPropertyRepo.CreateProperty(ctx, property); err != nil {
		return err
	}
	propertyAddr := &org_property_models.OrganizationPropertyAddress{
		ID:                     uuid.New(),
		OrganizationID:         input.OrgID,
		OrganizationPropertyID: property.ID,
		AddressLine:            input.AddressLine,
		City:                   input.City,
		State:                  input.State,
		Country:                input.Country,
		Pincode:                input.Pincode,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	if err := c.orgPropertyRepo.AddPropertyAddress(ctx, propertyAddr); err != nil {
		return err
	}

	return nil
}

func (r *organizationPropertyService) GetPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) (*org_property_contracts.OrganizationProperties, error) {
	properties, err := r.orgPropertyRepo.GetAllPropertiesByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	readers := make([]org_property_contracts.PropertiesReader, 0, len(properties))
	for _, p := range properties {

		address, err := r.orgPropertyRepo.GetPropertyAddress(ctx, p.ID)
		if err != nil {
			return nil, err
		}

		domain, err := r.orgPropertyRepo.GetPropertyDomain(ctx, p.ID)
		if err != nil {
			return nil, err
		}

		reader := org_property_contracts.PropertiesReader{
			ID:     p.ID,
			Name:   p.Name,
			Code:   p.Code,
			City:   "",
			State:  "",
			Domain: "",
		}

		if address != nil {
			reader.City = address.City
			reader.State = address.State
		}

		if domain != nil {
			reader.Domain = domain.Domain
		}

		readers = append(readers, reader)
	}

	return &org_property_contracts.OrganizationProperties{
		Properties: readers,
	}, nil
}
