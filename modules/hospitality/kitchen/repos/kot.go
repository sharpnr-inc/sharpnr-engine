package kitchenrepo

import (
	"context"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"
)

func (r *kitchenRepository) CreateKitchenTicket(ctx context.Context, ticket *kitchen_models.KitchenTicket) error {
	return r.db.WithContext(ctx).Create(ticket).Error
}

func (r *kitchenRepository) CreateKitchenTicketItem(ctx context.Context, item *kitchen_models.KitchenTicketItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *kitchenRepository) CreateKitchenTicketItemAddon(ctx context.Context, addon *kitchen_models.KitchenTicketItemAddon) error {
	return r.db.WithContext(ctx).Create(addon).Error
}
