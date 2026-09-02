package orgCostCenterDtos

import "github.com/google/uuid"

type CreateCostCenterRequestDto struct {
	Name        string `json:"name" validate:"required" binding:"required"`
	Description string `json:"description" validate:"required" binding:"required"`

	// Code is the short human key ("BLR-01"), unique per organization. It is
	// optional on the wire: when it is empty the service derives one from the
	// name, because the column is NOT NULL and every screen shows it.
	Code string `json:"code,omitempty"`

	Address string `json:"address" validate:"required" binding:"required"`
	City    string `json:"city" validate:"required" binding:"required"`
	State   string `json:"state" validate:"required" binding:"required"`
	Contry  string `json:"country" validate:"required" binding:"required"`
}

type CostCenterResponseDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Code        string    `json:"code"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Contry      string    `json:"country"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
