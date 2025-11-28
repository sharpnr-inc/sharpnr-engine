package kitchenrepo

import (
	"context"
	"errors"
	"fmt"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/******************************************************
 * This is the Repository of Kitchen                  *
 * and This section Contain all the operation for KOT *
******************************************************/

func (d *kitchenRepository) CreateKitchenTicketDBTx(ctx context.Context, ticket *hospitalitymodels.KitchenTicket) (*hospitalitymodels.KitchenTicket, error) {
	if err := d.db.WithContext(ctx).Create(&ticket).Error; err != nil {
		return nil, err
	}

	return ticket, nil
}

func (d *kitchenRepository) CreateKitchenTicketItemDBTx(ctx context.Context, item *hospitalitymodels.KitchenTicketItem) (*hospitalitymodels.KitchenTicketItem, error) {
	if err := d.db.WithContext(ctx).Create(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (d *kitchenRepository) GetKitchenTicketByIdDBTx(ctx context.Context, ticketId uuid.UUID) (*hospitalitymodels.KitchenTicket, error) {
	var ticket hospitalitymodels.KitchenTicket

	// Preload all needed relations
	err := d.db.WithContext(ctx).
		Model(&hospitalitymodels.KitchenTicket{}).
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
