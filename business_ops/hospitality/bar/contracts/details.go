package bar_contracts

import "github.com/google/uuid"

type BarDetailsReader struct {
	ID     uuid.UUID
	Name   string
	Status string
}
