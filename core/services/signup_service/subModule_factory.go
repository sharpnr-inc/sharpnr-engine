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

	// So we need to Create A Default Menu Group here
}

// case --> RESTAURANT
func (s *signupService) SetUpRestaurant(ctx context.Context, org *models.Organization, kitchen *hospitalitymodels.Kitchen) error {
	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,

		IsPureVeg:   true,
		OpeningTime: "09:00",
		ClosingTime: "21:00",
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
		OpeningTime:    "09:00",
		ClosingTime:    "21:00",
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
		OpeningTime:    "09:00",
		ClosingTime:    "21:00",
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
		Organization:   org,
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
	barFloor := &hospitalitymodels.BarFloor{
		ID:             uuid.New(),
		BarID:          defaultBar.ID,
		OrganizationID: defaultBar.OrganizationID,
		Organization:   org,
		Bar:            defaultBar,
		Name:           "Ground Floor",
		FloorNumber:    0,
	}

	if err := s.barRepo.CreateFloor(ctx, barFloor); err != nil {
		return fmt.Errorf("failed to create default floor")
	}

	// ➕ Create BAR Default MenuGroup
	barMenuGroup := &hospitalitymodels.MenuGroup{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "Bar Menu",
		Description:    "Default bar menu created during setup",
		Categories:     nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Type:           "BAR",
	}

	if err := s.menuRepo.CreateMenuGroup(ctx, barMenuGroup); err != nil {
		return fmt.Errorf("failed to create bar menu group: %w", err)
	}

	hotel := &hospitalitymodels.Hotel{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "__DEFAULT__HOTEL__NAME__",
		Organization:   org,
		NumberOfFloors: 0,
		NumberOfRooms:  0,
		Floors:         nil,
		IsKitchen:      true,
		KitchenID:      &kitchen.ID,
	}
	if err := s.hotelRepo.CreateHotel(ctx, hotel); err != nil {
		return fmt.Errorf("failed to create hotel for org %s: %w", org.ID, err)
	}

	// ➕ Create HOTEL Default MenuGroup
	hotelMenuGroup := &hospitalitymodels.MenuGroup{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "Room Service Menu",
		Description:    "Default hotel menu created during setup",
		Categories:     nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Type:           "HOTEL",
	}

	if err := s.menuRepo.CreateMenuGroup(ctx, hotelMenuGroup); err != nil {
		return fmt.Errorf("failed to create hotel menu group: %w", err)
	}

	defaultRestaurant := &hospitalitymodels.Restaurant{
		ID:             uuid.New(),
		Name:           "__DEFAULT__RES__NAME__",
		OrganizationID: org.ID,
		Organization:   org,
		IsPureVeg:      true,
		OpeningTime:    "09:00",
		ClosingTime:    "21:00",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		IsKitchen:      true,
		KitchenID:      &kitchen.ID,
	}
	if err := s.restaurantRepo.CreateRestaurant(ctx, defaultRestaurant); err != nil {
		return fmt.Errorf("failed to create restaurant for org %s: %w", org.ID, err)
	}

	defaultFloor := &hospitalitymodels.RestaurantFloor{
		ID:             uuid.New(),
		RestaurantID:   defaultRestaurant.ID,
		Restaurant:     defaultRestaurant,
		OrganizationID: org.ID,
		Organization:   org,
		Name:           "Ground Floor",
		Description:    "Ground floor or Restaurant",
		FloorNumber:    0,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.restaurantRepo.CreateOrAddFloorInRestaurant(ctx, defaultFloor); err != nil {
		return fmt.Errorf("failed to create restaurant default floor")
	}

	restaurantMenuGroup := &hospitalitymodels.MenuGroup{
		ID:             uuid.New(),
		OrganizationID: org.ID,
		Name:           "Restaurant Menu",
		Description:    "Default restaurant menu created during setup",
		Categories:     nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Type:           "RESTAURANT",
	}

	if err := s.menuRepo.CreateMenuGroup(ctx, restaurantMenuGroup); err != nil {
		return fmt.Errorf("failed to create restaurant menu group: %w", err)
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
