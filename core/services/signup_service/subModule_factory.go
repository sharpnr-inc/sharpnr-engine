package signupservice

import (
	"context"
	"fmt"
	"time"

	"sharpnr-engine/core/models"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	sharpnruirenderer "sharpnr-engine/sharpnr_ui_renderer"

	"github.com/google/uuid"
)

// --------------------------- SINGLE MODULE CASES ---------------------------

// case --> HOTEL
func (s *signupService) SetUpHotel(ctx context.Context, org *models.Organization, kitchen *hospitalitymodels.Kitchen) error {
	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "Default Hotel",

		NumberOfFloors: 0,
		NumberOfRooms:  0,
		Floors:         nil,
		IsKitchen:      true,
		KitchenID:      &kitchen.ID,
	}

	if err := s.hotelRepo.CreateHotel(ctx, hotel); err != nil {
		return fmt.Errorf("failed to create hotel for org %s: %w", org.ID, err)
	}

	newRendererID := uuid.New()
	ui_render, err := sharpnruirenderer.HotelRendererFactory(ctx, org.ID, org.OrgOperationType, newRendererID)
	if err != nil {
		return fmt.Errorf("failed to create hotel renderer: %w", err)
	}
	fmt.Println("✅ Renderer created successfully")

	rendererTree := &models.OrgUIRendererTree{
		ID:             newRendererID,
		UIRendererTree: ui_render,
		OrganizationID: org.ID,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.uiRenderRepo.CreateOrgUIRenderer(ctx, rendererTree); err != nil {
		return err
	}

	fmt.Printf("✅ Hotel created successfully for org: %s\n", org.ID)
	return nil
}

// case --> BAR
func (s *signupService) SetUpBar(ctx context.Context, org *models.Organization, kitchen *hospitalitymodels.Kitchen) error {
	defaultBar := &hospitalitymodels.Bar{
		ID:             uuid.New(),
		Name:           "__DEFAULT__BAR__NAME__",
		OrganizationID: org.ID,

		OpeningTime: time.Now(),
		ClosingTime: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsKitchen:   true,
		KitchenID:   &kitchen.ID,
	}

	if err := s.barRepo.CreateBar(ctx, defaultBar); err != nil {
		return fmt.Errorf("failed to create bar for org %s: %w", org.ID, err)
	}

	newRendererID := uuid.New()
	ui_render, err := sharpnruirenderer.BarRendererFactory(ctx, org.ID, org.OrgOperationType, newRendererID)
	if err != nil {
		return fmt.Errorf("failed to create bar renderer: %w", err)
	}
	fmt.Println("✅ Renderer created successfully")

	rendererTree := &models.OrgUIRendererTree{
		ID:             newRendererID,
		UIRendererTree: ui_render,
		OrganizationID: org.ID,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.uiRenderRepo.CreateOrgUIRenderer(ctx, rendererTree); err != nil {
		return err
	}

	fmt.Printf("✅ Bar created successfully for org: %s\n", org.ID)
	return nil
}

// case --> RESTAURANT
func (s *signupService) SetUpRestaurant(ctx context.Context, org *models.Organization, kitchen *hospitalitymodels.Kitchen) error {
	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,

		IsPureVeg:   true,
		OpeningTime: time.Now(),
		ClosingTime: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsKitchen:   true,
		KitchenID:   &kitchen.ID,
	}

	if err := s.restaurantRepo.CreateRestaurant(ctx, defaultRestaurant); err != nil {
		return fmt.Errorf("failed to create restaurant for org %s: %w", org.ID, err)
	}

	newRendererID := uuid.New()
	ui_render, err := sharpnruirenderer.RestaurantRendererFactory(ctx, org.ID, org.OrgOperationType, newRendererID)
	if err != nil {
		return fmt.Errorf("failed to create restaurant renderer: %w", err)
	}
	fmt.Println("✅ Renderer created successfully")

	rendererTree := &models.OrgUIRendererTree{
		ID:             newRendererID,
		UIRendererTree: ui_render,
		OrganizationID: org.ID,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.uiRenderRepo.CreateOrgUIRenderer(ctx, rendererTree); err != nil {
		return err
	}

	fmt.Printf("✅ Restaurant created successfully for org: %s\n", org.ID)
	return nil
}

// --------------------------- TWO MODULE COMBINATIONS ---------------------------

// case --> BAR + HOTEL
func (s *signupService) SetUpHotelBar(ctx context.Context, org *models.Organization, kitchen *hospitalitymodels.Kitchen) error {
	// create BAR
	defaultBar := &hospitalitymodels.Bar{
		ID:             uuid.New(),
		Name:           "__DEFAULT__BAR__NAME__",
		Floors:         nil,
		OrganizationID: org.ID,

		OpeningTime: time.Now(),
		ClosingTime: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsKitchen:   true,
		KitchenID:   &kitchen.ID,
	}
	if err := s.barRepo.CreateBar(ctx, defaultBar); err != nil {
		return fmt.Errorf("failed to create bar for org %s: %w", org.ID, err)
	}

	// create HOTEL
	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "__DEFAULT__HOTEL__NAME__",

		NumberOfFloors: 0,
		NumberOfRooms:  0,
		Floors:         nil,
		IsKitchen:      true,
		KitchenID:      &kitchen.ID,
	}
	if err := s.hotelRepo.CreateHotel(ctx, hotel); err != nil {
		return fmt.Errorf("failed to create hotel for org %s: %w", org.ID, err)
	}

	// single renderer for BAR + HOTEL
	newRendererID := uuid.New()
	ui_render, err := sharpnruirenderer.BarHotelRendererFactory(ctx, org.ID, org.OrgOperationType, newRendererID)
	if err != nil {
		return fmt.Errorf("failed to create BAR+HOTEL renderer: %w", err)
	}
	fmt.Println("✅ Renderer created successfully")

	rendererTree := &models.OrgUIRendererTree{
		ID:             newRendererID,
		UIRendererTree: ui_render,
		OrganizationID: org.ID,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.uiRenderRepo.CreateOrgUIRenderer(ctx, rendererTree); err != nil {
		return err
	}

	fmt.Printf("✅ Hotel + Bar created successfully for org: %s\n", org.ID)
	return nil
}

// case --> HOTEL + RESTAURANT
func (s *signupService) SetUpHotelRestaurant(ctx context.Context, org *models.Organization, kitchen *hospitalitymodels.Kitchen) error {
	// HOTEL
	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "__DEFAULT__HOTEL__NAME__",

		NumberOfFloors: 0,
		NumberOfRooms:  0,
		Floors:         nil,
		IsKitchen:      true,
		KitchenID:      &kitchen.ID,
	}
	if err := s.hotelRepo.CreateHotel(ctx, hotel); err != nil {
		return fmt.Errorf("failed to create hotel for org %s: %w", org.ID, err)
	}

	// RESTAURANT
	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,
		IsPureVeg:      true,
		OpeningTime:    time.Now(),
		ClosingTime:    time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		IsKitchen:      true,
		KitchenID:      &kitchen.ID,
	}
	if err := s.restaurantRepo.CreateRestaurant(ctx, defaultRestaurant); err != nil {
		return fmt.Errorf("failed to create restaurant for org %s: %w", org.ID, err)
	}

	newRendererID := uuid.New()
	ui_render, err := sharpnruirenderer.HotelRestaurantFactory(ctx, org.ID, org.OrgOperationType, newRendererID)
	if err != nil {
		return fmt.Errorf("failed to create HOTEL+RESTAURANT renderer: %w", err)
	}
	fmt.Println("✅ Renderer created successfully")

	rendererTree := &models.OrgUIRendererTree{
		ID:             newRendererID,
		UIRendererTree: ui_render,
		OrganizationID: org.ID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if err := s.uiRenderRepo.CreateOrgUIRenderer(ctx, rendererTree); err != nil {
		return err
	}

	fmt.Printf("✅ Hotel + Restaurant created successfully for org: %s\n", org.ID)
	return nil
}

// case --> BAR + RESTAURANT
func (s *signupService) SetUpBarRestaurant(ctx context.Context, org *models.Organization, kitchen *hospitalitymodels.Kitchen) error {
	defaultBar := &hospitalitymodels.Bar{
		ID:             uuid.New(),
		Name:           "__DEFAULT__BAR__NAME__",
		OrganizationID: org.ID,
		OpeningTime:    time.Now(),
		ClosingTime:    time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		IsKitchen:      true,
		KitchenID:      &kitchen.ID,
	}
	if err := s.barRepo.CreateBar(ctx, defaultBar); err != nil {
		return fmt.Errorf("failed to create bar for org %s: %w", org.ID, err)
	}

	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,
		IsPureVeg:      true,
		OpeningTime:    time.Now(),
		ClosingTime:    time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		IsKitchen:      true,
		KitchenID:      &kitchen.ID,
	}
	if err := s.restaurantRepo.CreateRestaurant(ctx, defaultRestaurant); err != nil {
		return fmt.Errorf("failed to create restaurant for org %s: %w", org.ID, err)
	}

	newRendererID := uuid.New()
	ui_render, err := sharpnruirenderer.BarRestaurantFactory(ctx, org.ID, org.OrgOperationType, newRendererID)
	if err != nil {
		return fmt.Errorf("failed to create BAR+RESTAURANT renderer: %w", err)
	}
	fmt.Println("✅ Renderer created successfully")

	rendererTree := &models.OrgUIRendererTree{
		ID:             newRendererID,
		UIRendererTree: ui_render,
		OrganizationID: org.ID,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.uiRenderRepo.CreateOrgUIRenderer(ctx, rendererTree); err != nil {
		return err
	}

	fmt.Printf("✅ Bar + Restaurant created successfully for org: %s\n", org.ID)
	return nil
}

// --------------------------- THREE MODULE COMBINATION ---------------------------

// case --> BAR + HOTEL + RESTAURANT
func (s *signupService) SetUpBarHotelRestaurant(ctx context.Context, org *models.Organization, kitchen *hospitalitymodels.Kitchen) error {
	defaultBar := &hospitalitymodels.Bar{
		ID:             uuid.New(),
		Name:           "__DEFAULT__BAR__NAME__",
		OrganizationID: org.ID,

		OpeningTime: time.Now(),
		ClosingTime: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsKitchen:   true,
		KitchenID:   &kitchen.ID,
	}
	if err := s.barRepo.CreateBar(ctx, defaultBar); err != nil {
		return fmt.Errorf("failed to create bar for org %s: %w", org.ID, err)
	}

	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "__DEFAULT__HOTEL__NAME__",

		NumberOfFloors: 0,
		NumberOfRooms:  0,
		Floors:         nil,
		IsKitchen:      true,
		KitchenID:      &kitchen.ID,
	}
	if err := s.hotelRepo.CreateHotel(ctx, hotel); err != nil {
		return fmt.Errorf("failed to create hotel for org %s: %w", org.ID, err)
	}

	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,

		IsPureVeg:   true,
		OpeningTime: time.Now(),
		ClosingTime: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsKitchen:   true,
		KitchenID:   &kitchen.ID,
	}
	if err := s.restaurantRepo.CreateRestaurant(ctx, defaultRestaurant); err != nil {
		return fmt.Errorf("failed to create restaurant for org %s: %w", org.ID, err)
	}

	newRendererID := uuid.New()
	ui_render, err := sharpnruirenderer.BarHotelRestaurantFactory(ctx, org.ID, org.OrgOperationType, newRendererID)
	if err != nil {
		return fmt.Errorf("failed to create BAR+HOTEL+RESTAURANT renderer: %w", err)
	}
	fmt.Println("✅ Renderer created successfully")

	rendererTree := &models.OrgUIRendererTree{
		ID:             newRendererID,
		UIRendererTree: ui_render,
		OrganizationID: org.ID,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.uiRenderRepo.CreateOrgUIRenderer(ctx, rendererTree); err != nil {
		return err
	}

	fmt.Printf("✅ Bar + Hotel + Restaurant created successfully for org: %s\n", org.ID)
	return nil
}
