package stockwaveEngineErrs

import "errors"

var (
	// ErrNoAccountFound is the org that never started stockwave. Every other
	// stockwave flow branches on exactly this, which is why it is a named error
	// rather than a bare gorm miss.
	ErrNoAccountFound = errors.New("no stockwave engine account found")

	// ErrAccountDisabled is the account that exists but was switched off. Kept
	// apart from "not found" all the way to the client because the fix differs
	// — start it vs. it was turned off.
	ErrAccountDisabled = errors.New("stockwave engine is disabled")

	// ErrStockwaveServiceAlreadyStarted answers a second start on an account
	// that is already live. Starting a switched-off account is not this error —
	// it re-activates the existing one.
	ErrStockwaveServiceAlreadyStarted = errors.New("stockwave service is already started for this organization")

	// ErrTermsNotAccepted guards the start flow: the caller has to actually
	// agree before the org gets a stockwave account.
	ErrTermsNotAccepted = errors.New("the terms and conditions must be accepted to start the stockwave service")
)
