// Package sharpCrmEngineErrs defines the sentinel errors returned by the
// SharpCRM module. Keeping them in one place lets the service signal a precise
// failure and the controllers map it to a status code with errors.Is, rather
// than matching on message text.
package sharpCrmEngineErrs

import "errors"

var (
	NoCustomerFoundOfthisID = errors.New("no customer found of this ID")

	/* --------------------------------- account -------------------------------- */

	// ErrCrmServiceNotInitialized is the org that never started SharpCRM. Every
	// other CRM flow branches on exactly this, which is why it is a named error
	// rather than a bare gorm miss.
	ErrCrmServiceNotInitialized = errors.New("sharpcrm service is not started, please initialize it first before using it")

	// ErrCrmAccountIsDisabled is the account that exists but was switched off.
	// Kept apart from "not initialized" all the way to the client because the
	// fix differs — start it vs. it was turned off.
	ErrCrmAccountIsDisabled = errors.New("sharpcrm account is disabled for the org, please contact support")

	// ErrCrmServiceAlreadyStarted answers a second start on an account that is
	// already live. Starting a switched-off account is not this error — it
	// re-activates the existing one.
	ErrCrmServiceAlreadyStarted = errors.New("sharpcrm service is already started for this organization")

	// ErrTermsNotAccepted guards the start flow: the caller has to actually
	// agree before the org gets a CRM account.
	ErrTermsNotAccepted = errors.New("the terms and conditions must be accepted to start the sharpcrm service")
)
