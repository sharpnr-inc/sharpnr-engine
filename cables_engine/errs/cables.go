package cablesEngineErrs

import "errors"

var (
	ErrCablesServiceNotInitialized = errors.New("cables service is not started, please initialize it first before using it")
	ErrNoChannelFound              = errors.New("no channel found for the org")
	ErrCablesAccountIsDisabled     = errors.New("cables account is disabled for the org, please contact support")
)
