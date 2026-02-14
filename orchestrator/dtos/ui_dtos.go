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
