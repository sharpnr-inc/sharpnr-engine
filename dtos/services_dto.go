package dto

type CreateServiceRequestDTO struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Duration    int     `json:"duration" binding:"required,gte=0"` // in minutes
}

type UpdateServiceRequestDTO struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Duration    *int     `json:"duration,omitempty"`
}

type CreateServiceResponseDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	Duration    int     `json:"duration"`
	CreatedAt   string  `json:"createdAt"`
}

type UpdateServiceResponseDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	Duration    int     `json:"duration"`
	UpdatedAt   string  `json:"updatedAt"`
}
