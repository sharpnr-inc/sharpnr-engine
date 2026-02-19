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

func (r *restaurantRepo) GetOrderItemVariantByOrderID(ctx context.Context, restaurantOrderID uuid.UUID) ([]*restaurant_models.RestaurantOrderItemsVariants, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var itemsVariant []*restaurant_models.RestaurantOrderItemsVariants

	err := db.WithContext(ctx).Where("restaurant_order_item_id = ?", restaurantOrderID).
		Order("created_at ASC").
		Find(&itemsVariant).
		Error
	if err != nil {
		return nil, err
	}

	return itemsVariant, nil
}

func (r *restaurantRepo) GetOrderItemAddOnByOrderID(ctx context.Context, restaurantOrderID uuid.UUID) ([]*restaurant_models.RestaurantOrderItemsAddOns, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var itemsAddon []*restaurant_models.RestaurantOrderItemsAddOns
	err := db.WithContext(ctx).Where("restaurant_order_item_id = ?", restaurantOrderID).
		Order("created_at ASC").
		Find(&itemsAddon).
		Error
	if err != nil {
		return nil, err
	}

	return itemsAddon, nil

}

/* ----------------------------------- <*> ---------------------------------- */

func (r *restaurantRepo) CreateOrder(ctx context.Context, order *restaurant_models.RestaurantOrders) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Persist the order to the database using the provided context
	err := db.WithContext(ctx).Create(order).Error
	if err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *restaurantRepo) CreateOrderItem(ctx context.Context, orderItem *restaurant_models.RestaurantOrderItems) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Insert the individual line item
	err := db.WithContext(ctx).Create(orderItem).Error
	if err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *restaurantRepo) CreateOrderItemVariant(ctx context.Context, orderVariant *restaurant_models.RestaurantOrderItemsVariants) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Insert the individual line item
	err := db.WithContext(ctx).Create(orderVariant).Error
	if err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */

func (r *restaurantRepo) CreateOrderItemAddon(ctx context.Context, orderAddon *restaurant_models.RestaurantOrderItemsAddOns) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Insert the individual line item
	err := db.WithContext(ctx).Create(orderAddon).Error
	if err != nil {
		return err
	}

	return nil
}

/* ---------------------------------- <***> --------------------------------- */
