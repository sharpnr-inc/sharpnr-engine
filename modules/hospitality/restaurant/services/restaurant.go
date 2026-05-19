package restaurantservices

import (
	"context"
	"errors"
	restaurant_constants "sharpnr-engine/modules/hospitality/restaurant/constants"
	restaurant_contracts "sharpnr-engine/modules/hospitality/restaurant/contracts"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *restaurantService) SetupDefaultRestaurant(ctx context.Context, orgID uuid.UUID, propertyID uuid.UUID) (restuarantID uuid.UUID, err error) {

	var defaultRestaurnID *uuid.UUID

	err = r.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		defRes := &restaurant_models.Restaurant{
			ID:             uuid.New(),
			Name:           "Default Restaurant",
			OrganizationID: orgID,
			PropertyID:     propertyID,
			CreatedAt:      time.Now(),
		}

		if err := r.restaurantRepo.CreateRestaurant(txCtx, defRes); err != nil {
			return err
		}

		if err := r.AddConfigurationStatus(txCtx, orgID, defRes.ID, restaurant_constants.RestaurantStatusCreated); err != nil {
			return err
		}
		defaultRestaurnID = &defRes.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return *defaultRestaurnID, nil

}

/* ----------------------------------- *** ---------------------------------- */

func (r *restaurantService) GetAllRestaurantOfProperty(
	ctx context.Context,
	propertyID uuid.UUID,
) ([]*restaurant_contracts.RestaurantDetails, error) {

	var result []*restaurant_contracts.RestaurantDetails

	err := r.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		restaurants, err := r.restaurantRepo.GetRestaurantByPropertyID(txCtx, propertyID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				result = []*restaurant_contracts.RestaurantDetails{}
				return nil
			}
			return err
		}

		result = make([]*restaurant_contracts.RestaurantDetails, 0, len(restaurants))

		for _, rs := range restaurants {

			resConfig, err := r.restaurantRepo.GetCurrentConfig(txCtx, rs.ID)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			details := &restaurant_contracts.RestaurantDetails{
				ID:         rs.ID,
				PropertyID: rs.PropertyID,
				Name:       rs.Name,
			}

			// fill configuration fields if present
			if resConfig != nil {
				details.Status = resConfig.IsActive
				details.IsConfigured = resConfig.IsConfigured
				details.ConfigurationStatus = resConfig.CurrentStatus
			}

			result = append(result, details)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

/* ---------------------------------- <**> ---------------------------------- */

func (r *restaurantService) GetRestaurantLayout(
	ctx context.Context,
	orgID uuid.UUID,
	restaurantID uuid.UUID,
) (layout restaurant_contracts.RestaurantLayout_Response_dto, err error) {

	err = r.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		floors, err := r.restaurantRepo.GetAllFloorOfRestaurant(txCtx, restaurantID)
		if err != nil {
			return err
		}

		var floorLayouts []restaurant_contracts.RestaurantFloorLayout_Response_dto

		for _, floor := range floors {

			tables, err := r.restaurantRepo.GetAllTablesOfFloor(txCtx, floor.ID)
			if err != nil {
				return err
			}

			var tableDTOs []restaurant_contracts.RestaurantTableLayout_Response_dto

			for _, table := range tables {

				// fetch dining session
				session, err := r.restaurantRepo.GetDiningSessionByTableID(txCtx, table.ID)

				var sessionID *uuid.UUID

				if err == nil && session != nil {
					sessionID = &session.ID
				} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}

				tableDTOs = append(tableDTOs, restaurant_contracts.RestaurantTableLayout_Response_dto{
					TableName:            table.Name,
					TableCode:            table.Code,
					FloorID:              table.FloorID.String(),
					TableID:              table.ID.String(),
					Status:               table.Status,
					Capacity:             table.Capacity,
					CurrentDineSessionID: sessionID,
				})
			}

			floorLayouts = append(floorLayouts, restaurant_contracts.RestaurantFloorLayout_Response_dto{
				Floor_Name:   floor.Name,
				Floor_Number: floor.FloorNumber,
				Tables:       tableDTOs,
			})
		}

		layout.Layout = floorLayouts
		return nil
	})

	if err != nil {
		return layout, err
	}

	return layout, nil
}

/* -------------------------------- **<***>** ------------------------------- */

func (s *restaurantService) GetRestaurantsOfOrg(
	ctx context.Context,
	orgID uuid.UUID,
) ([]restaurant_contracts.RestaurantDetails, error) {

	results := make([]restaurant_contracts.RestaurantDetails, 0)

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		restaurants, err := s.restaurantRepo.GetRestaurantByOrgID(txCtx, orgID)
		if err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				results = []restaurant_contracts.RestaurantDetails{}
				return nil
			}

			return err
		}

		results = make([]restaurant_contracts.RestaurantDetails, 0, len(restaurants))

		for _, r := range restaurants {

			resConfig, err := s.restaurantRepo.GetCurrentConfig(txCtx, r.ID)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			dto := restaurant_contracts.RestaurantDetails{
				ID:         r.ID,
				PropertyID: r.PropertyID,
				Name:       r.Name,
			}

			if resConfig != nil {
				dto.Status = resConfig.IsActive
				dto.IsConfigured = resConfig.IsConfigured
				dto.ConfigurationStatus = resConfig.CurrentStatus
			}

			results = append(results, dto)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}
