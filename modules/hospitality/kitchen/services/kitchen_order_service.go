package kitchenservices

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *kitchenService) GetOrdersList(
	ctx context.Context,
	orgId uuid.UUID,
) ([]interface{}, error) {

	// 1. Get Kitchen by Org
	kitchen, err := s.kitchenRepo.GetKitchenByOrgId(ctx, orgId)
	if err != nil {
		return nil, err
	}

	// 2. Get Tickets for the Kitchen
	tickets, err := s.kitchenRepo.GetKitchenTickets(ctx, kitchen.ID)
	if err != nil {
		return nil, err
	}

	// 3. Convert []KitchenTicket → []interface{}
	var response []interface{}
	for _, t := range tickets {
		response = append(response, t)
	}

	return response, nil
}

func (s *kitchenService) UpdateOrder(orderID string, status string) error {
	// TODO: Real DB update through Repo
	fmt.Printf("✅ Order [%s] updated to status: %s\n", orderID, status)
	return nil
}
