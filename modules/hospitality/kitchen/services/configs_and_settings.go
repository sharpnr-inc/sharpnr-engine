package kitchenservices

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	kitchen_contracts "sharpnr-engine/modules/hospitality/kitchen/contracts"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func (c *kitchenService) UpdateGeneralDetails(ctx context.Context, kitchenID uuid.UUID, req_dto kitchen_contracts.UpdateKitchenGeneralDetails_req_dto) error {
	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		// Initialize with mandatory fields
		updates := map[string]interface{}{
			"name":       req_dto.Name,
			"updated_at": time.Now(),
		}

		// Conditionally add description only if it contains data
		// This assumes that if it's empty, you don't want to clear it
		if req_dto.Description != "" {
			updates["description"] = req_dto.Description
		}

		return c.kitchenRepo.UpdateKitchenGeneralDetails(txCtx, kitchenID, updates)
	})
}

func (c *kitchenService) GetKitchenConfigs(ctx context.Context, kitchenID uuid.UUID) (*kitchen_contracts.KitchenConfigResponseDto, error) {
	resp := &kitchen_contracts.KitchenConfigResponseDto{}

	// We use a Mutex if we were writing to complex shared objects,
	// but since we are assigning to fields in a single struct,
	// it is safe as long as no two goroutines write to the same field.

	err := c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)
		g, gCtx := errgroup.WithContext(txCtx)

		g.Go(func() error {
			cfg, err := c.kitchenRepo.GetKitchenPrinterConfigByKitchenID(gCtx, kitchenID)
			if err != nil && err != gorm.ErrRecordNotFound {
				return err
			}
			if cfg != nil {
				resp.PrinterType = cfg.PrinterType
				resp.PrinterIP = cfg.IPAddress
				// Note: PrinterPort is not in your model, ensure you have
				// the right field or add it to the model.
			}
			return nil
		})

		g.Go(func() error {
			cfg, err := c.kitchenRepo.GetKitchenAutoAcceptOrdersByKitchenID(gCtx, kitchenID)
			if err != nil && err != gorm.ErrRecordNotFound {
				return err
			}
			if cfg != nil {
				resp.AutoAccept = cfg.IsAutoAcceptEnabled
			}
			return nil
		})

		g.Go(func() error {
			cfg, err := c.kitchenRepo.GetKitchenThrottlingConfigByKitchenID(gCtx, kitchenID)
			if err != nil && err != gorm.ErrRecordNotFound {
				return err
			}
			if cfg != nil {
				resp.Throttling.Enabled = cfg.IsEnabled
				resp.Throttling.MaxOrdersPerSlot = cfg.MaxOrdersPerSlot
				resp.Throttling.SlotDurationMins = cfg.SlotDurationMinutes
			}
			return nil
		})

		g.Go(func() error {
			hours, err := c.kitchenRepo.ListKitchenOperatingHoursByKitchenID(gCtx, kitchenID)
			if err != nil && err != gorm.ErrRecordNotFound {
				return err
			}

			for _, h := range hours {
				resp.OperatingHours = append(resp.OperatingHours, struct {
					DayOfWeek   int    `json:"day_of_week"`
					OpeningTime string `json:"opening_time"`
					ClosingTime string `json:"closing_time"`
				}{
					DayOfWeek:   h.DayOfWeek,
					OpeningTime: h.OpeningTime,
					ClosingTime: h.ClosingTime,
				})
			}
			return nil
		})

		return g.Wait()
	})

	return resp, err
}
