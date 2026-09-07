package stockwaveEngineErrs

import "errors"

var (
	ErrNoAccountFound  = errors.New("no stockwave engine account found")
	ErrAccountDisabled = errors.New("stockwave engine is disabled")
)
