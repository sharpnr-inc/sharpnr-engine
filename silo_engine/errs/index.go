// Package siloEngineErrs defines the sentinel errors returned by the Silo
// module. Keeping them in one place lets the service signal a precise failure
// and the controllers map it to a status code with errors.Is, rather than
// matching on message text.
package siloEngineErrs

import "errors"

var (
	// ErrSiloServiceNotInitialized is the org that never started Silo. Every
	// other Silo flow branches on exactly this, which is why it is a named error
	// rather than a bare gorm miss.
	ErrSiloServiceNotInitialized = errors.New("silo service is not started, please initialize it first before using it")

	// ErrSiloAccountIsDisabled is the account that exists but was switched off.
	// Kept apart from "not initialized" all the way to the client because the
	// fix differs — start it vs. it was turned off.
	ErrSiloAccountIsDisabled = errors.New("silo account is disabled for the org, please contact support")

	// ErrSiloServiceAlreadyStarted answers a second start on an account that is
	// already live. Starting a switched-off account is not this error — it
	// re-activates the existing one.
	ErrSiloServiceAlreadyStarted = errors.New("silo service is already started for this organization")

	// ErrTermsNotAccepted guards the start flow: the caller has to actually
	// agree before the org gets a Silo account.
	ErrTermsNotAccepted = errors.New("the terms and conditions must be accepted to start the silo service")
)
