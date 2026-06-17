package orchestrator_dtos

import (
	"github.com/google/uuid"
)

// --------------------------------------------------------------------------------------//

type PropertiesOfOrganization_ResDto struct {
	Properties []PropertiesReaderSlice `json:"properties"`
}

type PropertiesReaderSlice struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	AddressLine string    `json:"address_line"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Code        string    `json:"code"`
}

// --------------------------------------------------------------------------------------//

type PropertyDetails_res_dto struct {
	Name        string                   `json:"name"`
	Code        string                   `json:"code"`
	City        string                   `json:"city"`
	AddressLine string                   `json:"address_line"`
	State       string                   `json:"state"`
	Country     string                   `json:"country"`
	Modules     *[]ModulesDetails_dto    `json:"__modules__data__,omitempty"`
	Infra       *[]PropertyInfra_res_dto `json:"infra,omitempty"`
}

type ModulesDetails_dto struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Status       bool      `json:"status"`
	Module_Type  string    `json:"module_type"`
	IsConfigured bool      `json:"is_configured"`
}

type PropertyInfra_res_dto struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Status     bool      `json:"status"`
	Infra_Type string    `json:"infra_type"`
}

// --------------------------------------------------------------------------------------//

type CreatePropertyReq_dto struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	AddressLine string `json:"address_line"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Pincode     string `json:"pincode"`
}

/* ------------------------------------ . ----------------------------------- */

type ListAllInfraOfProperty_Res_dto struct {
	PropertyID      uuid.UUID `json:"property_id"`
	Infrastructures []*PropertyInfra_res_dto
}
