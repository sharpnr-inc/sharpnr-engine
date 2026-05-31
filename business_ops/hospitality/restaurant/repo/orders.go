package restaurantrepo

import (
	"context"
	"fmt"
	restaurant_models "sharpnr-engine/business_ops/hospitality/restaurant/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (r *restaurantRepo) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error {
	result := r.db.WithContext(ctx).
		Model(&restaurant_models.RestaurantOrders{}).
		Where("id = ?", orderID).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	// optional safety check
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
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

func (r *restaurantRepo) GetOrderSequenceNumber(ctx context.Context, restaurantID uuid.UUID, orgID uuid.UUID, nowTime time.Time) (string, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// 1. Define the start and end of the current day
	startOfDay := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var lastOrder restaurant_models.RestaurantOrders

	// 2. Find the latest order for today, scoped to this Org and Restaurant
	// We order by CreatedAt DESC to get the absolute latest one
	err := db.WithContext(ctx).
		Where("organization_id = ?", orgID).
		Where("restaurant_id = ?", restaurantID).
		Where("created_at >= ? AND created_at < ?", startOfDay, endOfDay).
		Order("created_at DESC").
		First(&lastOrder).Error

	// 3. Handle the result
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// No orders today yet, start at 1
			return "1", nil
		}
		return "", err
	}

	// 4. Increment the existing number
	currentNum, err := strconv.Atoi(lastOrder.OrderNumber)
	if err != nil {
		// Fallback if the OrderNumber was somehow not a number
		return "1", fmt.Errorf("invalid order number format in database: %v", err)
	}

	nextNum := currentNum + 1
	return strconv.Itoa(nextNum), nil
}
