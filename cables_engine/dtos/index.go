package cablesEngineDtos

type StartCablesServiceRequestsDto struct {
	AgreeToTermsAndCondtions bool `json:"agreeToTermsAndCondtions" binding:"required"`
	StartService             bool `json:"startService" binding:"required"`
}
type CablesServiceResponseDto struct {
	IsStarted bool `json:"isStarted"`
	IsActive  bool `json:"isActive"`
}
