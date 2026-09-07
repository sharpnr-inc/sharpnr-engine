// All the error for the hrms Engine
package hrmsEngineErrs

import "errors"

var (
	// ErrNoHRMSAccountFound is the org that never started HRMS. Every other
	// HRMS flow branches on exactly this, which is why it is a named error
	// rather than a bare gorm miss.
	ErrNoHRMSAccountFound = errors.New("no hrms account for the organization")

	// ErrAccountDisabled is the account that exists but was switched off. Kept
	// apart from "not found" all the way to the client because the fix differs
	// — start it vs. it was turned off.
	ErrAccountDisabled = errors.New("hrms account is disabled")

	// ErrHrmsServiceAlreadyStarted answers a second start on an account that is
	// already live. Starting a switched-off account is not this error — it
	// re-activates the existing one.
	ErrHrmsServiceAlreadyStarted = errors.New("hrms service is already started for this organization")

	// ErrTermsNotAccepted guards the start flow: the caller has to actually
	// agree before the org gets an HRMS account.
	ErrTermsNotAccepted = errors.New("the terms and conditions must be accepted to start the hrms service")
)
