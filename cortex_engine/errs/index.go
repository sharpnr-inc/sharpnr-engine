// Package cortexEngineErrs defines the sentinel errors returned by the Cortex
// module. Keeping them in one place lets the service signal a precise failure
// and the controllers map it to a status code with errors.Is, rather than
// matching on message text.
package cortexEngineErrs

import "errors"

var (
	// ErrNoCortexAccountFound is the org that never started Cortex. Every other
	// Cortex flow branches on exactly this, which is why it is a named error
	// rather than a bare gorm miss.
	ErrNoCortexAccountFound = errors.New("no cortex account found the organization")

	// ErrCortexAccountIsDisabled is the account that exists but was switched
	// off. Kept apart from "not found" all the way to the client because the fix
	// differs — start it vs. it was turned off.
	ErrCortexAccountIsDisabled = errors.New("cortex account is not active or disabled")

	// ErrCortexServiceAlreadyStarted answers a second start on an account that
	// is already live. Starting a switched-off account is not this error — it
	// re-activates the existing one.
	ErrCortexServiceAlreadyStarted = errors.New("cortex service is already started for this organization")

	// ErrTermsNotAccepted guards the start flow: the caller has to actually
	// agree before the org gets a Cortex account.
	ErrTermsNotAccepted = errors.New("the terms and conditions must be accepted to start the cortex service")
)
