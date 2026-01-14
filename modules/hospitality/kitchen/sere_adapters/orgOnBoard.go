package kitchen_sere_adapters

import (
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"
	"sharpnr-engine/sere"
	"time"

	"github.com/google/uuid"
)

func createDefaultKitchen(
	kitchenRepo kitchenrepo.KitchenRepositoryInterface,
) sere.WorkflowHandler {
	return func(w *sere.Workflow) error {

		orgId := w.Data["org_id"].(uuid.UUID)

		defaultKitchen := &kitchen_models.Kitchen{
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
