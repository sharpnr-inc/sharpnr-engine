package dto

type CreateProductRequestDTO struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
}

type UpdateProductRequestDTO struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Stock       *int     `json:"stock,omitempty"`
}

type CreateProductResponseDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CreatedAt   string  `json:"createdAt"`
}

type UpdateProductResponseDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	UpdatedAt   string  `json:"updatedAt"`
}
