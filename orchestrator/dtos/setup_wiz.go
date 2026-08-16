package orchestrator_dtos

import (
	org_setup_wizards_constants "sharpnr-engine/core/org_setup_wizard/constants"
	orgSetUpWizardErrs "sharpnr-engine/core/org_setup_wizard/errs"
	"strings"
	"time"

	"github.com/google/uuid"
)

/* ---------------------------------- <***> --------------------------------- */

// This DTO is Request DTO for the API that handles the organization property details step of the setup wizard. It captures all the necessary information about the organization's properties, such
// as its name, address, and geolocation. This information is crucial for setting up the organization's profile and ensuring accurate location-based services.
type SetupOrgPropertyReqDTO struct {
	StepID       string  `json:"step_id" binding:"required"`
	Skipped      bool    `json:"skipped"`
	Name         *string `json:"name,omitempty"`
	PropertyCode *string `json:"property_code,omitempty"`
	AddressLine  *string `json:"address_line,omitempty"`
	City         *string `json:"city,omitempty"`
	State        *string `json:"state,omitempty"`
	Country      *string `json:"country,omitempty"`
	Pincode      *string `json:"pincode,omitempty"`

	// Optional geolocation fields to capture the organization's location for location-based services and analytics.
	Longitude *float64 `json:"longitude,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	AccuracyM *int     `json:"accuracy_m,omitempty"`
}

// The property fields are optional on the wire so that the step can be skipped,
// which means they cannot be enforced by binding tags. They are required only
// when the step is actually being completed.
func (r *SetupOrgPropertyReqDTO) Validate() error {
	if r.Skipped {
		return nil
	}

	for _, field := range []*string{
		r.Name, r.AddressLine, r.City, r.State, r.Country, r.Pincode,
	} {
		if field == nil || strings.TrimSpace(*field) == "" {
			return orgSetUpWizardErrs.PropertyDetailsRequiredErr
		}
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

/* ---------------------------------- <***> --------------------------------- */

// This is the Request DTO for the API that handles the legal entity details step of the setup wizard.
// It captures all the necessary information about a legal entity, such as its name, address, and type. This
// information is crucial for setting up the organization's legal structure and ensuring compliance with regulatory requirements.
type SetupLegalEntityReqDTO struct {
	StepID      uuid.UUID `json:"step_id" binding:"required"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	AddressLine string    `json:"address_line"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Pincode     string    `json:"pincode"`
}

/* ---------------------------------- <***> --------------------------------- */

/* ---------------------------------- <***> --------------------------------- */

// This is Response DTO for the API that returns the overall progress of the setup wizard, including the status of each individual step.
// It provides a comprehensive view of where the organization stands in the setup
// process and what steps have been completed, are in progress, or are pending.
type SetupProgressResponseDTO struct {
	OverallStatus org_setup_wizards_constants.OverallSetupStatus `json:"overall_status"`
	CurrentStep   *org_setup_wizards_constants.SetupStepType     `json:"current_step,omitempty"`
	IsCompleted   bool                                           `json:"is_completed"`
	Steps         []SetupStepProgressResponseDTO                 `json:"steps"`
}

type SetupStepProgressResponseDTO struct {
	ID          uuid.UUID                                 `json:"id"`
	Step        org_setup_wizards_constants.SetupStepType `json:"step"`
	Status      org_setup_wizards_constants.SetupStatus   `json:"status"`
	Version     int                                       `json:"version"`
	StartedAt   time.Time                                 `json:"started_at"`
	CompletedAt *time.Time                                `json:"completed_at,omitempty"`
}

/* ---------------------------------- <***> --------------------------------- */

// This Dto for the Completion of the Super Admin of the Organization Super Admin is already Created during the Signup process but this step is for capturing additional details of the Super Admin
// which are not captured during the signup process. This includes fields like profile photo, secondary email, and employee ID which are important for the organization's internal records and //
// communication.
type SetupSuperAdminDetailsReqDTO struct {
	StepID         uuid.UUID `json:"step_id" binding:"required"`
	FirstName      string    `json:"first_name" binding:"required"`
	LastName       string    `json:"last_name" binding:"required"`
	SecondaryEmail string    `json:"secondary_email" binding:"required,email"`
	EmployeeID     *string   `json:"employee_id,omitempty"`
}

/* ---------------------------------- <***> --------------------------------- */

type SetupWizUpdateOrgNameRequestDto struct {
	StepID  uuid.UUID `json:"step_id" binding:"required"`
	OrgName string    `json:"org_name"`
}
