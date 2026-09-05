// Package folioEngineErrs defines the sentinel errors returned by the Folio
// module. Keeping them in one place lets the service signal a precise failure
// and the controllers map it to a status code with errors.Is, rather than
// matching on message text.
package folioEngineErrs

import "errors"

var (
	// ErrFolioServiceNotInitialized is the org that never started Folio. Every
	// other Folio flow branches on exactly this, which is why it is a named
	// error rather than a bare gorm miss.
	ErrFolioServiceNotInitialized = errors.New("folio service is not started, please initialize it first before using it")

	// ErrFolioAccountIsDisabled is the account that exists but was switched off.
	// Kept apart from "not initialized" all the way to the client because the
	// fix differs — start it vs. it was turned off.
	ErrFolioAccountIsDisabled = errors.New("folio account is disabled for the org, please contact support")

	// ErrFolioServiceAlreadyStarted answers a second start on an account that is
	// already live. Starting a switched-off account is not this error — it
	// re-activates the existing one.
	ErrFolioServiceAlreadyStarted = errors.New("folio service is already started for this organization")

	// ErrTermsNotAccepted guards the start flow: the caller has to actually
	// agree before the org gets a Folio account.
	ErrTermsNotAccepted = errors.New("the terms and conditions must be accepted to start the folio service")
)
