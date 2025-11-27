package kitchenrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
)

/******************************************************
 * This is the Repository of Kitchen                  *
 * and This section Contain all the operation for KOT *
******************************************************/

func (d *kitchenRepository) CreateKOT(ctx context.Context, kot *hospitalitymodels.KitchenTicket) {}

// This function will update the order status of item in the KOT
func (d *kitchenRepository) UpdateOrderStatusOfKOT(ctx context.Context) {}
