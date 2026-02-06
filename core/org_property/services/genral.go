package org_property_services

import (
	"context"
	"errors"
	property_constants "sharpnr-engine/core/org_property/constants"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	org_property_models "sharpnr-engine/core/org_property/models"
	"time"

	"github.com/google/uuid"
)

func (c *organizationPropertyService) CreatePropertyDefault(ctx context.Context, input org_property_contracts.CreateOrgPropertyInput) error {

	property := &org_property_models.OrganizationProperty{
		ID:             uuid.New(),
		OrganizationID: input.OrgID,
		Code:           input.Code,
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
			ID:           p.ID,
			Name:         p.Name,
			Code:         p.Code,
			City:         "",
			State:        "",
			Domain:       "",
			IsConfigured: p.IsConfigured,
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

func (r *organizationPropertyService) GetPropertyDetailsByID(
	ctx context.Context,
	orgID uuid.UUID,
	propertyID uuid.UUID,
) (*org_property_contracts.PropertyDetailsReader, error) {

	pr, err := r.orgPropertyRepo.GetPropertyByID(ctx, propertyID)
	if err != nil {
		return nil, err
	}
	// Safety check for pr
	if pr == nil {
		return nil, errors.New("property not found")
	}

	if pr.OrganizationID != orgID {
		return nil, errors.New("invalid property")
	}

	prDomain, err := r.orgPropertyRepo.GetPropertyDomain(ctx, propertyID)
	if err != nil {
		return nil, err
	}

	prLocation, err := r.orgPropertyRepo.GetPropertyAddress(ctx, propertyID)
	if err != nil {
		return nil, err
	}

	// If address is mandatory, return error. If optional, use empty strings.
	if prLocation == nil {
		return nil, errors.New("property address record is missing")
	}

	res := &org_property_contracts.PropertyDetailsReader{
		Domain:      prDomain.Domain, // Use the safe variable
		Name:        pr.Name,
		Code:        pr.Code,
		AddressLine: prLocation.AddressLine,
		City:        prLocation.City,
		State:       prLocation.State,
		Country:     prLocation.Country,
		Pincode:     prLocation.Pincode,
	}

	return res, nil
}

func (r *organizationPropertyService) AddDomain(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID, Domain property_constants.PropertyDomainTypes) error {

	pr, err := r.orgPropertyRepo.GetPropertyByID(ctx, propertyID)
	if err != nil {
		return err
	}

	if pr.ID != propertyID {
		return errors.New("invalid resources")
	}

	domain := org_property_models.OrganizationPropertyDomain{
		ID:                     uuid.New(),
		OrganizationPropertyID: propertyID,
		Domain:                 Domain,
		Status:                 "NOT_CONFIGURED",
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	}

	if err := r.orgPropertyRepo.AddDomainOfProperty(ctx, &domain); err != nil {
		return err
	}

	return nil
}
