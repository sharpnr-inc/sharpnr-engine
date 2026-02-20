package kitchenservices

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	kitchen_contracts "sharpnr-engine/modules/hospitality/kitchen/contracts"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *kitchenService) CreateStagingTicket(ctx context.Context, req kitchen_contracts.CreateKOT_req_dto) error {

	return s.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// Here we create the Staging Ticket with Default
		// values and Send to the kitchen
		stagingTicket := &kitchen_models.KitchenTicketStaging{
			ID:             uuid.New(),
			OrganizationID: req.OrganizationID,
			PropertyID:     req.PropertyID,
			KitchenID:      req.KitchenID,
			OrderID:        req.OrderID,
			CreatedAt:      time.Now(),
		}

		if err := s.stagingRepo.CreateKitchenTicketStaging(txCtx, stagingTicket); err != nil {
			return err
		}

		for _, itm := range req.Items {

			orderItems := &kitchen_models.KitchenTicketItemStaging{
				ID:                     uuid.New(),
				KitchenTicketStagingID: stagingTicket.ID,
				MenuItemID:             itm.MenuItemID,
				MenuItemName:           itm.MenuItemName,
				Quantity:               itm.Quantity,
				Remarks:                itm.Remarks,
				CreatedAt:              time.Now(),
			}

			if err := s.stagingRepo.CreateKitchenTicketItemStaging(txCtx, orderItems); err != nil {
				return err
			}

			for _, vr := range itm.Variants {
				orderItemsVariants := &kitchen_models.KitchenTicketItemVariantStaging{
					ID:                     uuid.New(),
					KitchenTicketStagingID: stagingTicket.ID,
					VariantID:              vr.ID,
					VariantName:            vr.Name,
					CreatedAt:              time.Now(),
				}

				if err := s.stagingRepo.CreateKitchenTicketItemVariantStaging(txCtx, orderItemsVariants); err != nil {
					return err
				}
			}

		}

		return nil
	})
}
