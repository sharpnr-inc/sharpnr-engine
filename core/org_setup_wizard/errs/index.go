// Package orgSetUpWizardErrs defines the errors used across the organization
// setup wizard. It provides a centralized location for error definitions, making
// it easier to manage and maintain error handling throughout the wizard flow.
package orgSetUpWizardErrs

import "errors"

var (
	InvalidStepIDErr     = errors.New("invalid stepID")
	InvalidStepStatusErr = errors.New("invalid step status for this operation")
	InvalidStepErr       = errors.New("invalid step for this operation")
	StepDoesntExistErr   = errors.New("step does not exist for this organization")

	NilSetupStepErr    = errors.New("organization setup is nil")
	EmptyStepStatusErr = errors.New("step status cannot be empty")

	OrgNotVerifiedErr = errors.New("organization must be verified before running the setup wizard")

	PropertyDetailsRequiredErr = errors.New("property details are required when the step is not skipped")
)
