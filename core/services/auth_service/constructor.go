package authservice

import "sharpnr-engine/core/interfaces"

type loginService struct {
	orgRepo interfaces.OrganizationRepository // interface, not concrete struct
}
