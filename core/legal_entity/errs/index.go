package legalEntityErrs

import "errors"

var (
	NoLegalEntityFoundForOrg = errors.New("no legal entity found for the organanization")
)
