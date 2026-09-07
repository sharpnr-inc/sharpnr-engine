// Package accountingEngineErrs defines the sentinel errors returned by the
// accounting engine's account layer. Keeping them in one place lets the service
// signal a precise failure and the controllers map it to a status code with
// errors.Is, rather than matching on message text.
package accountingEngineErrs

import "errors"

var (
	// ErrAccountingServiceNotInitialized is the org that never started the
	// accounting service. Every accounting-owned flow branches on exactly this,
	// which is why it is a named error rather than a bare gorm miss.
	ErrAccountingServiceNotInitialized = errors.New("accounting service is not started, please initialize it first before using it")

	// ErrAccountingAccountIsDisabled is the account that exists but was switched
	// off. Kept apart from "not initialized" all the way to the client because
	// the fix differs — start it vs. it was turned off.
	ErrAccountingAccountIsDisabled = errors.New("accounting account is disabled for the org, please contact support")

	// ErrAccountingServiceAlreadyStarted answers a second start on an account
	// that is already live. Starting a switched-off account is not this error —
	// it re-activates the existing one.
	ErrAccountingServiceAlreadyStarted = errors.New("accounting service is already started for this organization")

	// ErrTermsNotAccepted guards the start flow: the caller has to actually
	// agree before the org gets an accounting account.
	ErrTermsNotAccepted = errors.New("the terms and conditions must be accepted to start the accounting service")
)
