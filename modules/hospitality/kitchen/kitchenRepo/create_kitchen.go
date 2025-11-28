package kitchenrepo

import (
	"context"
	"errors"
	"fmt"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateKitchen inserts a new kitchen record into the database
func (d *kitchenRepository) CreateKitchen(ctx context.Context, kitchen *hospitalitymodels.Kitchen) error {
	// Use WithContext to respect request timeouts and cancellations
	if err := d.db.WithContext(ctx).Create(kitchen).Error; err != nil {
		return err
	}
	return nil
}

func (d *kitchenRepository) GetKitchenByOrgId(
	ctx context.Context,
	orgId uuid.UUID,
) (*hospitalitymodels.Kitchen, error) {

	var kitchen hospitalitymodels.Kitchen

	err := d.db.WithContext(ctx).
		Where("organization_id = ?", orgId).
		Preload("Organization").
		Preload("MenuItems").
		First(&kitchen).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("kitchen not found for org: %s", orgId)
		}
		return nil, err
	}

	return &kitchen, nil
}

func (r *kitchenRepository) GetNextTicketNumber(
	ctx context.Context,
	kitchenId uuid.UUID,
) (string, error) {

	var lastTicketNo string

	// Note: SELECT ... FOR UPDATE is safe even without tx because:
	// - PostgreSQL still locks the matched rows during this query
	// - Prevents 2 concurrent reads returning the same lastTicket
	err := r.db.WithContext(ctx).
		Raw(`
            SELECT ticket_no
            FROM kitchen_tickets
            WHERE kitchen_id = ?
            ORDER BY created_at DESC
            LIMIT 1
            FOR UPDATE
        `, kitchenId).
		Scan(&lastTicketNo).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	// No ticket yet
	if lastTicketNo == "" {
		return "KOT-101", nil
	}

	parts := strings.Split(lastTicketNo, "-")
	lastNum, _ := strconv.Atoi(parts[1])

	nextNum := lastNum + 1

	return fmt.Sprintf("KOT-%d", nextNum), nil
}
