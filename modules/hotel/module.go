package hotel

import (
	"fmt"
	"github.com/sharpnr-inc/sharpnr-engine/core/models"
	"github.com/sharpnr-inc/sharpnr-engine/modules/base"
)

// hotelModule implements the base.ModuleInitializer interface for Hotel organizations.
type hotelModule struct {
	// Add any dependencies needed for hotel initialization here (e.g., hotel-specific repository)
}

// NewHotelModule creates a new instance of the hotel module initializer.
func NewHotelModule() base.ModuleInitializer {
	return &hotelModule{}
}

// InitializeOrganization performs hotel-specific setup for a newly registered organization.
func (hm *hotelModule) InitializeOrganization(org *models.Organization) error {
	fmt.Printf("Initializing Hotel module for organization: %s (UUID: %s)\n", org.Name, org.SharpnrOrgUUID)

	// TODO: Implement actual hotel-specific initialization logic here.
	// This might include:
	// - Creating default room types (e.g., Standard, Deluxe, Suite)
	// - Setting up initial hotel configurations (e.g., check-in/check-out times)
	// - Populating default services or amenities
	// - Creating initial departments specific to hotels (e.g., Front Desk, Housekeeping)

	// Example: Create a default room type (conceptual, requires actual implementation)
	// roomTypeRepo.CreateDefaultRoomType(org.SharpnrOrgUUID, "Standard Room")

	return nil
}