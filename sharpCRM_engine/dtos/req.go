package sharpCrmEngineDtos

// StartCRMServiceRequestDto provisions the org's one SharpCRM account.
//
// Neither field carries binding:"required": on a bool that tag rejects `false`
// as "missing", which would turn a deliberate decline into a binding error
// instead of the ErrTermsNotAccepted the client can actually act on. Both are
// checked in the service.
type StartCRMServiceRequestDto struct {
	Agreed_terms_and_conditions bool `json:"agreedTermsAndConditions"`
	StartService                bool `json:"startService"`
}
