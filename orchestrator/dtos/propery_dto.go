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
	Name        string `json:"name"`
	Code        string `json:"code"`
	City        string `json:"city"`
	AddressLine string `json:"address_line"`
	State       string `json:"state"`
	Country     string `json:"country"`
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
