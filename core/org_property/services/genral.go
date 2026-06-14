package org_property_services

import (
	"context"
	"errors"
	org_property_contracts "sharpnr-engine/core/org_property/contracts"
	org_property_models "sharpnr-engine/core/org_property/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c *organizationPropertyService) CreatePropertyDefault(ctx context.Context, input org_property_contracts.CreateOrgPropertyInput) error {

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)
		now := time.Now()

		property := &org_property_models.OrganizationProperty{
			ID:             uuid.New(),
			OrganizationID: input.OrgID,
			Code:           input.Code,
			Name:           input.Name,
			CreatedAt:      now,
			UpdatedAt:      now,
		}
		if err := c.orgPropertyRepo.CreateProperty(txCtx, property); err != nil {
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
			CreatedAt:              now,
			UpdatedAt:              now,
		}

		if err := c.orgPropertyRepo.AddPropertyAddress(txCtx, propertyAddr); err != nil {
			return err
		}

		return nil
	})

}

/* ---------------------------------- <***> --------------------------------- */

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

		readers = append(readers, reader)
	}

	return &org_property_contracts.OrganizationProperties{
		Properties: readers,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

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

	prLocation, err := r.orgPropertyRepo.GetPropertyAddress(ctx, propertyID)
	if err != nil {
		return nil, err
	}

	// If address is mandatory, return error. If optional, use empty strings.
	if prLocation == nil {
		return nil, errors.New("property address record is missing")
	}

	res := &org_property_contracts.PropertyDetailsReader{
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

/* ---------------------------------- <***> --------------------------------- */
