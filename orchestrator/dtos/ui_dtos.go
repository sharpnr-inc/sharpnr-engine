package orchestrator_dtos

import "github.com/google/uuid"

// DomainConfig represents modules + infra inside a domain
type DomainConfig struct {
	Modules []string `json:"Modules"`
	Infras  []string `json:"Infras"`
}

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

/* ---------------------------------- <***> --------------------------------- */

type AllKitchenResponseDto struct {
	ID       uuid.UUID
	Name     string
	Property RestaurantProperty
}
