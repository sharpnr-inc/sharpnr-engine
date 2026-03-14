package kitchenservices

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	kitchen_contracts "sharpnr-engine/modules/hospitality/kitchen/contracts"
	"time"

	"github.com/google/uuid"
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

func (c *kitchenService) GetKitchenConfigs(
	ctx context.Context,
	kitchenID uuid.UUID,
) (*kitchen_contracts.KitchenConfigResponseDto, error) {

	resp := &kitchen_contracts.KitchenConfigResponseDto{}

	err := c.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		// 1️⃣ Get kitchen basic info
		kitchen, err := c.kitchenRepo.GetKitchenById(txCtx, kitchenID)
		if err != nil {
			return err
		}

		resp.IsActive = kitchen.IsActive

		// 2️⃣ Printer Config
		printerCfg, err := c.kitchenRepo.GetKitchenPrinterConfigByKitchenID(txCtx, kitchenID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if printerCfg != nil {
			resp.PrinterType = printerCfg.PrinterType
			resp.PrinterIP = printerCfg.IPAddress
		}

		// 3️⃣ Auto Accept Config
		autoCfg, err := c.kitchenRepo.GetKitchenAutoAcceptOrdersByKitchenID(txCtx, kitchenID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if autoCfg != nil {
			resp.AutoAccept = autoCfg.IsAutoAcceptEnabled
		}

		// 4️⃣ Throttling Config
		throttleCfg, err := c.kitchenRepo.GetKitchenThrottlingConfigByKitchenID(txCtx, kitchenID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if throttleCfg != nil {
			resp.Throttling.Enabled = throttleCfg.IsEnabled
			resp.Throttling.MaxOrdersPerSlot = throttleCfg.MaxOrdersPerSlot
			resp.Throttling.SlotDurationMins = throttleCfg.SlotDurationMinutes
		}

		// 5️⃣ Operating Hours
		hours, err := c.kitchenRepo.ListKitchenOperatingHoursByKitchenID(txCtx, kitchenID)
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

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *kitchenService) ActivateDeactivate(ctx context.Context, orgID, emplID, kitchenID uuid.UUID, kitchenSwitch bool) error {

	return c.txManager.WithTx(ctx, func(tx *gorm.DB) error {
		txCtx := dbtx.InjectTx(ctx, tx)

		if err := c.kitchenRepo.SwitchKitchenOnOff(txCtx, kitchenID, kitchenSwitch); err != nil {
			return err
		}

		return nil
	})
}
