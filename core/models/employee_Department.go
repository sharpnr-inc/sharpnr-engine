package models

import "github.com/google/uuid"

// this is the Table for the Employee and Department

type Employee_Department struct {
	ID           uuid.UUID
	Deparment_ID string
	Employee_ID  string
}
