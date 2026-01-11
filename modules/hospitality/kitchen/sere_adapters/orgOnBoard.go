package kitchen_sere_adapters

import (
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/sere"
	"time"

	"github.com/google/uuid"
)

func createDefaultKitchen(
	kitchenRepo kitchenrepo.KitchenRepositoryInterface,
) sere.WorkflowHandler {
	return func(w *sere.Workflow) error {

		orgId := w.Data["org_id"].(uuid.UUID)

		defaultKitchen := &hospitalitymodels.Kitchen{
			ID:             uuid.New(),
			Name:           "Default system Kitchen",
			Description:    "some thing",
			OrganizationID: orgId,
			Type:           "SYSTEM",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := kitchenRepo.CreateKitchen(w.Ctx, defaultKitchen); err != nil {
			return err
		}

		w.Data["kitchen_id"] = defaultKitchen.ID
		return nil

	}
}
