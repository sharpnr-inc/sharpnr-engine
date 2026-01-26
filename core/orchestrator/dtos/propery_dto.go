package orchestrator_dtos

import "github.com/google/uuid"

type PropertiesOfOrganization_ResDto struct {
	Properties []PropertiesReaderSlice
}

type PropertiesReaderSlice struct {
	ID     uuid.UUID
	Name   string
	City   string
	State  string
	Domain string
	Code   string
}

type PropertyDetails_res_dto struct {
	Name        string
	Code        string
	City        string
	AddressLine string
	State       string
	Country     string
	Domain      *string
	Modules     *HospitalityPropertyModules
}

type HospitalityPropertyModules struct {
	HotelDetails *[]HotelDetails_dto
}

type HotelDetails_dto struct {
	ID     uuid.UUID
	Name   string
	Status bool
}
