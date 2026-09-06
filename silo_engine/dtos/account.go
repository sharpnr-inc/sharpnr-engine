package siloEngineDtos

import "github.com/google/uuid"

// StartSiloServiceRequestDto provisions the org's one Silo account.
//
// Neither field carries binding:"required": on a bool that tag rejects `false`
// as "missing", which would turn a deliberate decline into a binding error
// instead of the ErrTermsNotAccepted the client can actually act on. Both are
// checked in the service.
type StartSiloServiceRequestDto struct {
	AgreeTermsAndConditions bool `json:"agreeTermsAndConditions"`
	StartService            bool `json:"startService"`
}

// SiloAccountStatusResponse is the read the client makes before it knows
// whether the org has Silo at all.
//
// An organization that never started the service answers serviceStarted=false
// with the other two fields null, rather than a 404 — a missing account is
// information here, not a failure. IsActive and ID are pointers so "not
// started" and "started but switched off" never collapse into the same JSON.
type SiloAccountStatusResponse struct {
	ServiceStarted bool       `json:"serviceStarted"`
	IsActive       *bool      `json:"isActive"`
	ID             *uuid.UUID `json:"id"`
}
