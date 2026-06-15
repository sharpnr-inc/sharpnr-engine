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
			AddressLine:    input.AddressLine,
			City:           input.City,
			CityCode:       input.Code,
			State:          input.State,
			Country:        input.Country,
			Pincode:        input.Pincode,
			CreatedAt:      now,
			UpdatedAt:      now,
		}
		if err := c.orgPropertyRepo.CreateProperty(txCtx, property); err != nil {
			return err
		}

		return nil
	})

}

/* ---------------------------------- <***> --------------------------------- */

func (r *organizationPropertyService) GetPropertiesOfOrg(ctx context.Context, orgID uuid.UUID) ([]org_property_contracts.PropertiesReader, error) {

	properties, err := r.orgPropertyRepo.GetAllPropertiesByOrgID(ctx, orgID)
	if err != nil {
		return nil, err
	}

	// Initialize the slice to hold the mapped results
	// Pre-allocating the capacity with len(properties) improves performance
	readers := make([]org_property_contracts.PropertiesReader, 0, len(properties))

	for _, p := range properties {
		reader := org_property_contracts.PropertiesReader{
			ID:          p.ID,
			Name:        p.Name,
			Code:        p.Code,
			City:        p.City,
			State:       p.State,
			AddressLine: p.AddressLine,
			Country:     p.Country,
		}

		// Append each contract reader to the slice
		readers = append(readers, reader)
	}

	// Return the slice
	return readers, nil
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

	res := &org_property_contracts.PropertyDetailsReader{
		Name:        pr.Name,
		Code:        pr.Code,
		AddressLine: pr.AddressLine,
		City:        pr.City,
		State:       pr.State,
		Country:     pr.Country,
		Pincode:     pr.Pincode,
	}

	return res, nil
}

/* ---------------------------------- <***> --------------------------------- */
