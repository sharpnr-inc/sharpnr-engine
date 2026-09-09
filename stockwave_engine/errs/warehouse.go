package stockwaveEngineErrs

import "errors"

var (
	ErrNoWareHouseFound = errors.New("no warehouse found for the organization")
)
