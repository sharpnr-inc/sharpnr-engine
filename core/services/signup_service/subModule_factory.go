package signupservice

import (
	"context"
	"fmt"
	"time"

	"sharpnr-engine/core/models"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

// --------------------------- SINGLE MODULE CASES ---------------------------

// case --> HOTEL
func (s *signupService) SetUpHotel(ctx context.Context, org *models.Organization) error {
	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "__DEFAULT__HOTEL__NAME__",
		Organization:   *org,
		NumberOfFloors: 0,
		NumberOfRooms:  0,
		Floors:         nil,
	}

	if err := s.hotelRepo.CreateHotel(ctx, hotel); err != nil {
		return fmt.Errorf("failed to create hotel for org %s: %w", org.ID, err)
	}

	fmt.Printf("✅ Hotel created successfully for org: %s\n", org.ID)
	return nil
}

// case --> BAR
func (s *signupService) SetUpBar(ctx context.Context, org *models.Organization) error {
	defaultBar := &hospitalitymodels.Bar{
		ID:             uuid.New(),
		Name:           "__DEFAULT__BAR__NAME__",
		OrganizationID: org.ID,
		Organization:   *org,
		OpeningTime:    time.Now(),
		ClosingTime:    time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.barRepo.CreateBar(ctx, defaultBar); err != nil {
		return fmt.Errorf("failed to create bar for org %s: %w", org.ID, err)
	}

	fmt.Printf("✅ Bar created successfully for org: %s\n", org.ID)
	return nil
}

// case --> RESTAURANT
func (s *signupService) SetUpRestaurant(ctx context.Context, org *models.Organization) error {
	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,
		Organization:   *org,
		IsPureVeg:      true,
		OpeningTime:    time.Now(),
		ClosingTime:    time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.restaurantRepo.CreateRestaurant(ctx, defaultRestaurant); err != nil {
		return fmt.Errorf("failed to create restaurant for org %s: %w", org.ID, err)
	}

	fmt.Printf("✅ Restaurant created successfully for org: %s\n", org.ID)
	return nil
}

// --------------------------- TWO MODULE COMBINATIONS ---------------------------

// case --> BAR + HOTEL
func (s *signupService) SetUpHotelBar(ctx context.Context, org *models.Organization) error {
	if err := s.SetUpBar(ctx, org); err != nil {
		return fmt.Errorf("failed to create bar in BAR_HOTEL setup: %w", err)
	}

	if err := s.SetUpHotel(ctx, org); err != nil {
		return fmt.Errorf("failed to create hotel in BAR_HOTEL setup: %w", err)
	}

	fmt.Printf("✅ Hotel + Bar created successfully for org: %s\n", org.ID)
	return nil
}

// case --> HOTEL + RESTAURANT
func (s *signupService) SetUpHotelRestaurant(ctx context.Context, org *models.Organization) error {
	if err := s.SetUpHotel(ctx, org); err != nil {
		return fmt.Errorf("failed to create hotel in HOTEL_RESTAURANT setup: %w", err)
	}

	if err := s.SetUpRestaurant(ctx, org); err != nil {
		return fmt.Errorf("failed to create restaurant in HOTEL_RESTAURANT setup: %w", err)
	}

	fmt.Printf("✅ Hotel + Restaurant created successfully for org: %s\n", org.ID)
	return nil
}

// case --> BAR + RESTAURANT
func (s *signupService) SetUpBarRestaurant(ctx context.Context, org *models.Organization) error {
	if err := s.SetUpBar(ctx, org); err != nil {
		return fmt.Errorf("failed to create bar in BAR_RESTAURANT setup: %w", err)
	}

	if err := s.SetUpRestaurant(ctx, org); err != nil {
		return fmt.Errorf("failed to create restaurant in BAR_RESTAURANT setup: %w", err)
	}

	fmt.Printf("✅ Bar + Restaurant created successfully for org: %s\n", org.ID)
	return nil
}

// --------------------------- THREE MODULE COMBINATION ---------------------------

// case --> BAR + HOTEL + RESTAURANT
func (s *signupService) SetUpBarHotelRestaurant(ctx context.Context, org *models.Organization) error {
	if err := s.SetUpBar(ctx, org); err != nil {
		return fmt.Errorf("failed to create bar in BAR_HOTEL_RESTAURANT setup: %w", err)
	}

	if err := s.SetUpHotel(ctx, org); err != nil {
		return fmt.Errorf("failed to create hotel in BAR_HOTEL_RESTAURANT setup: %w", err)
	}

	if err := s.SetUpRestaurant(ctx, org); err != nil {
		return fmt.Errorf("failed to create restaurant in BAR_HOTEL_RESTAURANT setup: %w", err)
	}

	fmt.Printf("✅ Bar + Hotel + Restaurant created successfully for org: %s\n", org.ID)
	return nil
}
