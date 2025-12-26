package acc_organization

import (
	"log/slog"
	"sharpnr-engine/core/core_modules_events"
	hospitality_acc_setup "sharpnr-engine/modules/accounting/accountingSetup/hospitalityAccSetup"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/sere"
)

func contains(arr []string, v string) bool {
	for _, s := range arr {
		if s == v {
			return true
		}
	}
	return false
}

func extractHospitalitySubModules(modules []string) []string {
	sub := []string{}
	for _, m := range modules {
		switch m {
		case "BAR", "HOTEL", "RESTAURANT":
			sub = append(sub, m)
		}
	}
	return sub
}

type OrganizationCreatedEventHandler struct {
	accRepo accounting_repo.AccountingRepositoryInterface
}

func NewOrganizationCreatedEventHandler(
	accRepo accounting_repo.AccountingRepositoryInterface,
) *OrganizationCreatedEventHandler {
	return &OrganizationCreatedEventHandler{
		accRepo: accRepo,
	}
}

func (h *OrganizationCreatedEventHandler) Handle(e sere.Event) error {

	event := e.Payload.(core_modules_events.OrganizationCreated)
	ctx := e.Ctx

	logger := slog.With(
		"handler", "OrganizationCreatedEventHandler",
		"org_id", event.OrganizationID,
	)

	logger.Info("organization created event received",
		"modules", event.Modules,
	)

	// ------------------------------------------------
	// Check if hospitality module is enabled
	// ------------------------------------------------
	// if !contains(event.OrganizationType, "HOSPITALITY") {
	// 	logger.Info("hospitality module not enabled, skipping accounting setup")
	// 	return nil
	// }

	subModules := extractHospitalitySubModules(event.Modules)

	if len(subModules) == 0 {
		logger.Warn("hospitality enabled but no submodules found, skipping accounting setup")
		return nil
	}

	logger.Info("hospitality accounting setup starting",
		"sub_modules", subModules,
	)

	// ------------------------------------------------
	// Build setup context
	// ------------------------------------------------
	setupCtx := hospitality_acc_setup.AccSetupContext{
		OrganizationID:   event.OrganizationID,
		OrganizationName: event.OrganizationName,
		OrganizationType: event.OrganizationType,
		Modules:          event.Modules,
		AccountingRepo:   h.accRepo,
	}

	steps := hospitality_acc_setup.HospitalityAccountingSetupFactory(subModules)

	// ------------------------------------------------
	// Execute setup steps
	// ------------------------------------------------
	for _, step := range steps {
		logger.Info("executing accounting setup step",
			"step", step.Name(),
		)

		if err := step.Execute(ctx, setupCtx); err != nil {
			logger.Error("accounting setup step failed",
				"step", step.Name(),
				"error", err,
			)
			return err
		}

		logger.Info("accounting setup step completed",
			"step", step.Name(),
		)
	}

	logger.Info("hospitality accounting setup completed successfully")

	return nil
}
