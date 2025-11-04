package signupservice

import (
	"context"
	"fmt"
	"sharpnr-engine/core/models"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"time"

	"github.com/google/uuid"
)

// --------------------------- SINGLE MODULE CASES ---------------------------

// case --> HOTEL
func (s *signupService) CreateHotel(ctx context.Context, org *models.Organization) error {
	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "__DEFAULT__HOTEL__NAME",
		Organization:   *org,
		NumberOfFloors: 0,
		NumberOfRooms:  0,
		Floors:         nil,
	}

	hoteldata, err := s.hotelRepo.CreateHotel(ctx, hotel)
	if err != nil {
		fmt.Println(hoteldata)
		return nil
	}

	return nil
}

// case --> BAR
func (s *signupService) CreateBar(ctx context.Context, org *models.Organization) {
	defaultBar := &hospitalitymodels.Bar{
		ID:             uuid.New(),
		Name:           "__DEFAULT__BAR__NAME__",
		OrganizationID: org.ID,
		Organization:   *org,
		OpeningTime:    time.Time{},
		ClosingTime:    time.Time{},
		Tables:         nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	// bar := s.barRepo.CreateBar()
	fmt.Println(defaultBar)

}

// case --> RESTAURANT
func (s *signupService) CreateRestaurant(ctx context.Context, org *models.Organization) {
	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,
		Organization:   *org,
		IsPureVeg:      true,
		OpeningTime:    time.Time{},
		ClosingTime:    time.Time{},
		Tables:         nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	fmt.Println(defaultRestaurant)
}

// --------------------------- TWO MODULE COMBINATIONS ---------------------------

// case --> HOTEL_BAR
func (s *signupService) CreateHotelBar(ctx context.Context, org *models.Organization) error {

	defaultBar := &hospitalitymodels.Bar{
		ID:             uuid.New(),
		Name:           "__DEFAULT__BAR__NAME__",
		OrganizationID: org.ID,
		Organization:   *org,
		OpeningTime:    time.Time{},
		ClosingTime:    time.Time{},
		Tables:         nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	fmt.Println(defaultBar)
	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "__DEFAULT__HOTEL__NAME",
		Organization:   *org,
		NumberOfFloors: 0,
		NumberOfRooms:  0,
		Floors:         nil,
	}

	hoteldata, err := s.hotelRepo.CreateHotel(ctx, hotel)
	if err != nil {
		fmt.Println(hoteldata)
		return nil
	}

	return nil

}

// case --> HOTEL_RESTAURANT
func (s *signupService) CreateHotelRestaurant(ctx context.Context, org *models.Organization) {
	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "__DEFAULT__HOTEL__NAME",
		Organization:   *org,
		NumberOfFloors: 0,
		NumberOfRooms:  0,
		Floors:         nil,
	}
	fmt.Println(hotel)

	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,
		Organization:   *org,
		IsPureVeg:      true,
		OpeningTime:    time.Time{},
		ClosingTime:    time.Time{},
		Tables:         nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	fmt.Println(defaultRestaurant)

}

// case --> BAR_RESTAURANT
func (s *signupService) CreateBarRestaurant(ctx context.Context, org *models.Organization) {
	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,
		Organization:   *org,
		IsPureVeg:      true,
		OpeningTime:    time.Time{},
		ClosingTime:    time.Time{},
		Tables:         nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	fmt.Println(defaultRestaurant)

	defaultBar := &hospitalitymodels.Bar{
		ID:             uuid.New(),
		Name:           "__DEFAULT__BAR__NAME__",
		OrganizationID: org.ID,
		Organization:   *org,
		OpeningTime:    time.Time{},
		ClosingTime:    time.Time{},
		Tables:         nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	fmt.Println(defaultBar)

}

// --------------------------- THREE MODULE COMBINATION ---------------------------

// case --> BAR_HOTEL_RESTAURANT
func (s *signupService) CreateBarHotelRestaurant(ctx context.Context, org *models.Organization) {
	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,
		Organization:   *org,
		IsPureVeg:      true,
		OpeningTime:    time.Time{},
		ClosingTime:    time.Time{},
		Tables:         nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	fmt.Println(defaultRestaurant)

	defaultBar := &hospitalitymodels.Bar{
		ID:             uuid.New(),
		Name:           "__DEFAULT__BAR__NAME__",
		OrganizationID: org.ID,
		Organization:   *org,
		OpeningTime:    time.Time{},
		ClosingTime:    time.Time{},
		Tables:         nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	fmt.Println(defaultBar)

	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "__DEFAULT__HOTEL__NAME",
		Organization:   *org,
		NumberOfFloors: 0,
		NumberOfRooms:  0,
		Floors:         nil,
	}
	fmt.Println(hotel)
}
