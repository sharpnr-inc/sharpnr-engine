package inventory_engine_services

import (
	"context"
	"fmt"
	inventory_constants "sharpnr-engine/service_engines/inventory_engine/constants"
	inventory_engine_contracts "sharpnr-engine/service_engines/inventory_engine/contracts"
	inventory_engine_models "sharpnr-engine/service_engines/inventory_engine/models"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (s *inventoryEngine) AddItemsInInventory(
	ctx context.Context,
	input inventory_engine_contracts.AddItemsInInventoryInput,
) error {
	now := time.Now()

	for _, line := range input.Lines {
		if err := s.processInventoryLine(ctx, input.EmployeeID, input.SupplierID, line, now); err != nil {
			return fmt.Errorf("failed on line (desc: %s): %w", line.Description, err)
		}
	}

	return nil
}

func (s *inventoryEngine) processInventoryLine(
	ctx context.Context,
	emplID uuid.UUID,
	supplierID uuid.UUID,
	line inventory_engine_contracts.AddItemsInInventoryLineInput,
	now time.Time,
) error {

	var (
		itemID         uuid.UUID
		quantityBefore float64
	)

	// ── Existing item: just bump quantity ─────────────────────────────────────
	if line.InventoryItemID != nil {
		itemID = *line.InventoryItemID

		existing, err := s.repo.GetInventoryItemByID(ctx, itemID)
		if err != nil {
			return fmt.Errorf("item not found (%s): %w", itemID, err)
		}
		quantityBefore = existing.Quantity

		if err := s.repo.AdjustInventoryItemQuantity(ctx, itemID, line.Qty); err != nil {
			return fmt.Errorf("adjust quantity failed: %w", err)
		}

		// ── New item: create it with opening qty ──────────────────────────────────
	} else {
		quantityBefore = 0
		category := line.Category
		if category == "" {
			category = "OTHER"
		}

		newItem := &inventory_engine_models.InventoryItem{
			ID:            uuid.New(),
			InventoryID:   line.InventoryID,
			SupplierID:    &supplierID,
			COGSAccountID: line.COGSAccountID,
			Name:          line.Description,
			SKU:           generateSKUFromDescription(line.Description),
			Unit:          line.Unit,
			Category:      category,
			Quantity:      line.Qty,
			Status:        "ACTIVE",
		}

		if err := s.repo.CreateInventoryItem(ctx, newItem); err != nil {
			return fmt.Errorf("create item failed: %w", err)
		}

		itemID = newItem.ID
	}

	// ── Price: close old → open new ───────────────────────────────────────────
	if old, err := s.repo.GetActiveInventoryItemPriceBySupplier(
		ctx, itemID, supplierID, inventory_constants.InventortItemPricePurchasePrice,
	); err == nil && old != nil {
		if err := s.repo.DeactivateInventoryItemPrice(ctx, old.ID, now); err != nil {
			return fmt.Errorf("deactivate old price failed: %w", err)
		}
	}

	if err := s.repo.CreateInventoryItemPrice(ctx, &inventory_engine_models.InventoryItemPrice{
		ID:              uuid.New(),
		InventoryItemID: itemID,
		SupplierID:      &supplierID,
		PriceType:       inventory_constants.InventortItemPricePurchasePrice,
		Price:           line.UnitPrice,
		Currency:        "INR",
		EffectiveFrom:   now,
	}); err != nil {
		return fmt.Errorf("create price failed: %w", err)
	}

	// ── History: PURCHASE_RECEIVED ────────────────────────────────────────────
	if err := s.repo.CreateInventoryItemHistory(ctx, &inventory_engine_models.InventoryItemHistory{
		ID:              uuid.New(),
		InventoryItemID: itemID,
		ChangeType:      "PURCHASE_RECEIVED",
		ReferenceType:   "PO",
		QuantityBefore:  quantityBefore,
		QuantityChange:  line.Qty,
		QuantityAfter:   quantityBefore + line.Qty,
		UnitCost:        line.UnitPrice,
		TotalCost:       line.Qty * line.UnitPrice,
		PerformedBy:     emplID,
		CreatedAt:       now,
	}); err != nil {
		return fmt.Errorf("create history failed: %w", err)
	}

	return nil
}

// Generates a slug-style SKU: "Chicken Breast" → "CHICKEN-BREAST-<short uuid>"
func generateSKUFromDescription(description string) string {
	slug := strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(description), " ", "-"))
	short := uuid.New().String()[:8] // collision buffer
	return fmt.Sprintf("%s-%s", slug, short)
}
