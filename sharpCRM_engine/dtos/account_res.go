package sharpCrmEngineDtos

import "github.com/google/uuid"

// SharpCrmAccountStatusResponse is the read the client makes before it knows
// whether the org has SharpCRM at all.
//
// An organization that never started the service answers serviceStarted=false
// with the other two fields null, rather than a 404 — a missing account is
// information here, not a failure. IsActive and ID are pointers so "not started"
// and "started but switched off" never collapse into the same JSON.
type SharpCrmAccountStatusResponse struct {
	ServiceStarted bool       `json:"serviceStarted"`
	IsActive       *bool      `json:"isActive"`
	ID             *uuid.UUID `json:"id"`
}
