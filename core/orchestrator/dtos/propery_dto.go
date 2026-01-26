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
