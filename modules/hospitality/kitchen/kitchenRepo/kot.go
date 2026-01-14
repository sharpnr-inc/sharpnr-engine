package kitchenrepo

import (
	"context"
	"errors"
	"fmt"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/******************************************************
 * This is the Repository of Kitchen                  *
 * and This section Contain all the operation for KOT *
******************************************************/

func (d *kitchenRepository) CreateKitchenTicketDBTx(ctx context.Context, ticket *kitchen_models.KitchenTicket) (*kitchen_models.KitchenTicket, error) {
	if err := d.db.WithContext(ctx).Create(&ticket).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}

func (d *kitchenRepository) CreateKitchenTicketItemDBTx(ctx context.Context, item *kitchen_models.KitchenTicketItem) (*kitchen_models.KitchenTicketItem, error) {
	if err := d.db.WithContext(ctx).Create(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (d *kitchenRepository) GetKitchenTicketByIdDBTx(ctx context.Context, ticketId uuid.UUID) (*kitchen_models.KitchenTicket, error) {
	var ticket kitchen_models.KitchenTicket

	// Preload all needed relations
	err := d.db.WithContext(ctx).
		Model(&kitchen_models.KitchenTicket{}).
		Where("id = ?", ticketId).
		Preload("Organization").
		Preload("Kitchen").
		Preload("Items").
		First(&ticket).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("kitchen ticket not found")
		}
		return nil, err
	}

	return &ticket, nil
}

//------------------------------------------------------------------------------------------------------------//

func (d *kitchenRepository) GetKitchenTickets(
	ctx context.Context,
	kitchenId uuid.UUID,
) ([]kitchen_models.KitchenTicket, error) {

	var tickets []kitchen_models.KitchenTicket

	err := d.db.WithContext(ctx).
		Preload("Items").
		Preload("Kitchen").
		Preload("Organization").
		Where("kitchen_id = ?", kitchenId).
		Order("created_at DESC").
		Find(&tickets).Error

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (d *kitchenRepository) CreateKitchenTicketItemAddonDBTx(
	ctx context.Context,
	addon *kitchen_models.KitchenTicketItemAddon,
) (*kitchen_models.KitchenTicketItemAddon, error) {

	if err := d.db.WithContext(ctx).Create(&addon).Error; err != nil {
		return nil, err
	}

	return addon, nil
}
