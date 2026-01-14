package bar_sere_adapters

import (
	hospitality_org_factory_repository "sharpnr-engine/core/hospitality-org-factory/repository"
	bar_models "sharpnr-engine/modules/hospitality/bar/models"
	barrepo "sharpnr-engine/modules/hospitality/bar/repo"
	"sharpnr-engine/sere"
	"time"

	"github.com/google/uuid"
)

func defaultBarHandler(
	barRepo barrepo.BarRepository,
	hospitalityOrgFactoryRepository hospitality_org_factory_repository.HospitalityOrgFactoryRepoInterface,
) sere.WorkflowHandler {
	return func(wf *sere.Workflow) error {

		// ------------------------------------------------
		// 0️⃣ Check if HOTEL is selected in request
		// ------------------------------------------------
		modules := wf.Data["sector_modules"].([]string)

		hasBar := false
		for _, m := range modules {
			if m == "BAR" {
				hasBar = true
				break
			}
		}

		if !hasBar {
			return nil // 🚫 HOTEL not requested
		}

		// ------------------------------------------------
		// 1️⃣ Preconditions
		// ------------------------------------------------
		orgID := wf.Data["org_id"].(uuid.UUID)
		kitchenID := wf.Data["kitchen_id"].(uuid.UUID)

		defaultBar := &bar_models.Bar{
			ID:             uuid.New(),
			Name:           "Default System Created",
			OrganizationID: orgID,
			IsKitchen:      true,
			KitchenID:      (*uuid.UUID)(&kitchenID),
			CreatedAt:      time.Now(),
		}
		if err := barRepo.CreateBar(wf.Ctx, defaultBar); err != nil {
			return err
		}

		return nil
	}
}
