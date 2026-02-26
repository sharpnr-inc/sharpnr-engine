package orchestrator_dtos

import "github.com/google/uuid"

type PropertyOfResturants struct {
	Properties []PropertiesOfOrg
}

type PropertiesOfOrg struct {
	PropertyID   uuid.UUID
	PropertyName string
	Restaurents  []uuid.UUID
}

// DomainConfig represents modules + infra inside a domain
type DomainConfig struct {
	Modules []string `json:"Modules"`
	Infras  []string `json:"Infras"`
}

/*
Domains structure:

[

	{
	  "Hospitality": {
	    "Modules": [...],
	    "Infras": [...]
	  }
	}

]
*/
type OrgLauncherResponse struct {
	Domains []map[string]DomainConfig `json:"Domains"`
}

type AllRestaurantsResponseDto struct {
	ID           uuid.UUID
	Name         string
	OrgID        uuid.UUID
	IsActive     bool
	IsConfigured bool
	Property     RestaurantProperty
}

type RestaurantProperty struct {
	ID          uuid.UUID
	Name        string
	AddressLine string
	City        string
	State       string
	Country     string
}
