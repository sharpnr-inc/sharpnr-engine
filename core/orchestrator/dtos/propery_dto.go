package orchestrator_dtos

import (
	property_constants "sharpnr-engine/core/org_property/constants"

	"github.com/google/uuid"
)

type PropertiesOfOrganization_ResDto struct {
	Properties []PropertiesReaderSlice `jason:"properties"`
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

type PropertyDetails_res_dto struct {
	Name        string                                  `json:"name"`
	Code        string                                  `json:"code"`
	City        string                                  `json:"city"`
	AddressLine string                                  `json:"address_line"`
	State       string                                  `json:"state"`
	Country     string                                  `json:"country"`
	Domain      *property_constants.PropertyDomainTypes `json:"domain,omitempty"`
	// Modules will only show up if it's not nil
	Modules *HospitalityPropertyModules `json:"modules,omitempty"`
}

type HospitalityPropertyModules struct {
	HotelDetails   *[]HotelDetails_dto   `json:"hotel_details,omitempty"`
	KitchenDetails *[]KitchenDetails_dto `json:"kitchen_details,omitempty"`
}

type HotelDetails_dto struct {
	ID     uuid.UUID
	Name   string
	Status bool
}

type KitchenDetails_dto struct {
	ID     uuid.UUID
	Name   string
	Status bool
}

type AddDomainToPropertyReq_DTO struct {
	Domain property_constants.PropertyDomainTypes `json:"domain" binding:"required"`
}

type PropertyModuleConfigurationDTO struct {
	PropertyID uuid.UUID                              `json:"property_id"`
	Domain     property_constants.PropertyDomainTypes `json:"domain"`
	Modules    []ModuleConfiguration                  `json:"modules"`
}

type ModuleConfiguration struct {
	Name           string `json:"name"`
	KitchenCapable bool   `json:"kitchen_capable"`
	KitchenEnabled bool   `json:"kitchen_enabled"`
}

type CreatePropertyReq_dto struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	AddressLine string `json:"address_line"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Pincode     string `json:"pincode"`
}
