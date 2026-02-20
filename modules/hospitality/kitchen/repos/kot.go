package kitchenrepo

import (
	"context"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"
	"time"

	"github.com/google/uuid"
)

func (r *kitchenRepository) CreateKitchenTicket(ctx context.Context, ticket *kitchen_models.KitchenTicket) error {
	return r.db.WithContext(ctx).Create(ticket).Error
}

func (r *kitchenRepository) CreateKitchenTicketItem(ctx context.Context, item *kitchen_models.KitchenTicketItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *kitchenRepository) CreateKitchenTicketItemVariant(ctx context.Context, item *kitchen_models.KitchenTicketItemVariant) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *kitchenRepository) CreateKitchenTicketItemAddon(ctx context.Context, addon *kitchen_models.KitchenTicketItemAddon) error {
	return r.db.WithContext(ctx).Create(addon).Error
}

func (r *kitchenRepository) GetKitchenTicketsByKitchenID(ctx context.Context, kitchenID uuid.UUID) ([]*kitchen_models.KitchenTicket, error) {
	var tickets []*kitchen_models.KitchenTicket

	// We use WithContext to ensure transaction/tracing support
	// We order by CreatedAt ASC so the kitchen sees the oldest orders first (FIFO)
	err := r.db.WithContext(ctx).
		Table("hospitality.kitchen_ticket").
		Where("kitchen_id = ?", kitchenID).
		Order("created_at ASC").
		Find(&tickets).
		Error

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

// Fetches all items belonging to a specific Kitchen Ticket
func (r *kitchenRepository) GetKitchenTicketItemsByTicketID(ctx context.Context, kitchenTicketID uuid.UUID) ([]*kitchen_models.KitchenTicketItem, error) {
	var items []*kitchen_models.KitchenTicketItem
	err := r.db.WithContext(ctx).
		Where("kitchen_ticket_id = ?", kitchenTicketID).
		Find(&items).Error
	return items, err
}

// Fetches all addons belonging to a specific Kitchen Ticket Item
func (r *kitchenRepository) GetKitchenTicketItemAddOn(ctx context.Context, kitchenTicketItemID uuid.UUID) ([]*kitchen_models.KitchenTicketItemAddon, error) {
	var addons []*kitchen_models.KitchenTicketItemAddon
	err := r.db.WithContext(ctx).
		Where("kitchen_ticket_item_id = ?", kitchenTicketItemID).
		Find(&addons).Error
	return addons, err
}

func (r *kitchenRepository) GetKitchenTicketItemVariant(ctx context.Context, kitchenTicketItemID uuid.UUID) (*kitchen_models.KitchenTicketItemVariant, error) {
	var variant *kitchen_models.KitchenTicketItemVariant
	err := r.db.WithContext(ctx).
		Where("kitchen_ticket_item_id = ?", kitchenTicketItemID).
		Find(&variant).Error
	return variant, err
}

func (r *kitchenRepository) GetKOTByID(ctx context.Context, kotID uuid.UUID) (*kitchen_models.KitchenTicket, error) {
	var ticket kitchen_models.KitchenTicket
	err := r.db.WithContext(ctx).
		Where("id = ?", kotID).
		First(&ticket).Error

	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *kitchenRepository) UpdateKotStatus(ctx context.Context, kotID uuid.UUID, status string) error {
	updates := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}

	// Logic: If status is COMPLETED, set the CompletedAt timestamp
	if status == "COMPLETED" || status == "SERVED" {
		now := time.Now()
		updates["completed_at"] = &now
	}

	return r.db.WithContext(ctx).
		Model(&kitchen_models.KitchenTicket{}).
		Where("id = ?", kotID).
		Updates(updates).Error
}
