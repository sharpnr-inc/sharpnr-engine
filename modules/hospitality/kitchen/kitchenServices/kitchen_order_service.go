package kitchenservices

import "fmt"

func (s *kitchenService) GetOrders() ([]interface{}, error) {
	// TODO: Replace with actual query to HospitalityOrder/KitchenTicket tables
	orders := []interface{}{
		map[string]interface{}{
			"order_id":   "1",
			"item":       "Paneer Butter Masala",
			"status":     "PENDING",
			"kitchen_id": "default",
		},
		map[string]interface{}{
			"order_id":   "2",
			"item":       "Fried Rice",
			"status":     "IN_PROGRESS",
			"kitchen_id": "default",
		},
	}
	return orders, nil
}

func (s *kitchenService) UpdateOrder(orderID string, status string) error {
	// TODO: Real DB update through Repo
	fmt.Printf("✅ Order [%s] updated to status: %s\n", orderID, status)
	return nil
}
