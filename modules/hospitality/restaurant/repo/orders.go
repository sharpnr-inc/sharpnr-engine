package restaurantrepo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	restaurant_models "sharpnr-engine/modules/hospitality/restaurant/models"

	"github.com/google/uuid"
)

func (r *restaurantRepo) GetOrdersByDiningSessionID(ctx context.Context, sessionID uuid.UUID) ([]*restaurant_models.RestaurantOrders, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var orders []*restaurant_models.RestaurantOrders

	// We filter by the foreign key RestaurantDiningSessionID
	// We order by CreatedAt so the waiter sees orders in the sequence they were placed
	err := db.WithContext(ctx).
		Where("restaurant_dining_session_id = ?", sessionID).
		Order("created_at ASC").
		Find(&orders).
		Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

/* ----------------------------------- <*> ---------------------------------- */

func (r *restaurantRepo) GetOrderItemsByOrderID(ctx context.Context, restaurantOrderID uuid.UUID) ([]*restaurant_models.RestaurantOrderItems, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var items []*restaurant_models.RestaurantOrderItems

	// Filter by the parent Order ID
	// We order by CreatedAt to ensure items appear in the order they were requested
	err := db.WithContext(ctx).
		Where("restaurant_order_id = ?", restaurantOrderID).
		Order("created_at ASC").
		Find(&items).
		Error

	if err != nil {
		return nil, err
	}

	return items, nil
}

/* ----------------------------------- <*> ---------------------------------- */
