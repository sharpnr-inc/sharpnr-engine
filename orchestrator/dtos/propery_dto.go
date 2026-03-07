package orchestrator_dtos

import (
	property_constants "sharpnr-engine/core/org_property/constants"

	"github.com/google/uuid"
)

// --------------------------------------------------------------------------------------//

type PropertiesOfOrganization_ResDto struct {
	Properties []PropertiesReaderSlice `json:"properties"`
}

type PropertiesReaderSlice struct {
	ID           uuid.UUID                              `json:"id"`
	Name         string                                 `json:"name"`
	City         string                                 `json:"city"`
	State        string                                 `json:"state"`
	Domain       property_constants.PropertyDomainTypes `json:"domain"`
	IsConfigured bool                                   `json:"is_configured"`
	Code         string                                 `json:"code"`
}

// --------------------------------------------------------------------------------------//

type PropertyDetails_res_dto struct {
	Name        string                                  `json:"name"`
	Code        string                                  `json:"code"`
	City        string                                  `json:"city"`
	AddressLine string                                  `json:"address_line"`
	State       string                                  `json:"state"`
	Country     string                                  `json:"country"`
	Domain      *property_constants.PropertyDomainTypes `json:"domain,omitempty"`
	Modules     *[]ModulesDetails_dto                   `json:"__modules__data__,omitempty"`
	Infra       *[]PropertyInfra_res_dto                `json:"infra,omitempty"`
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

type AddDomainToPropertyReq_DTO struct {
	Domain property_constants.PropertyDomainTypes `json:"domain" binding:"required"`
}

// --------------------------------------------------------------------------------------//

type PropertyModuleConfigurationDTO struct {
	PropertyID     uuid.UUID                              `json:"property_id"`
	Domain         property_constants.PropertyDomainTypes `json:"domain"`
	Modules        *[]string                              `json:"modules,omitempty"`
	Infrastructure *[]PropertyInfraConfiguration          `json:"infrastructure,omitempty"`
}

type PropertyInfraConfiguration struct {
	Infra_Name string `json:"infra_name"`
	InfraQty   int    `json:"qty"`
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

// --------------------------------------------------------------------------------------//

type ConfigureHotel_Request_dto struct {
	PropertyID  uuid.UUID                                     `json:"propertyId"`
	HotelID     uuid.UUID                                     `json:"hotelId"`
	GenDetails  ConfigurationPropertyHotel_GenDetails_Req_dto `json:"genDetails"`
	HotelLayout []HotelLayoutSetup_Req_dto                    `json:"hotelLayout"`
}

type ConfigurationPropertyHotel_GenDetails_Req_dto struct {
	HotelDisplayName string `json:"hotelDisplayName"`
	Is24X7           bool   `json:"is24x7"`
	OpeningTime      string `json:"openingTime"`
	ClosingTime      string `json:"closingTime"`
}

type HotelLayoutSetup_Req_dto struct {
	FloorNumber   int    `json:"floorNumber"`
	FloorName     string `json:"floorName"`
	NumberOfRooms int    `json:"numberOfRooms"`
}

/* ------------------------- ModuleOperationsResults ------------------------ */

type ModuleOperationResult struct {
	Created []string `json:"created"`
	Skipped []string `json:"skipped"`
	Message string   `json:"message"`
}

/* ------------------------------------ . ----------------------------------- */

type ListAllInfraOfProperty_Res_dto struct {
	PropertyID      uuid.UUID `json:"property_id"`
	Infrastructures []*PropertyInfra_res_dto
}
