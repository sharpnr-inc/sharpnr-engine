package signupservice

import (
	"context"
	"fmt"
	"sharpnr-engine/core/models"
)

func (s *signupService) CreateHotel(ctx context.Context, org *models.Organization) {
	fmt.Println("This is Create Hotel")
}
