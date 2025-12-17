package accounting_sere_handlers

import (
	"sharpnr-engine/modules/accounting/accounting_models"
	"sharpnr-engine/modules/accounting/accounting_repo"
	"sharpnr-engine/sere"
	"sharpnr-engine/sere/events"
	"time"

	"github.com/google/uuid"
)

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

	event := e.Payload.(events.OrganizationCreated)
	ctx := e.Ctx
	// // ✅ Idempotency guard (VERY IMPORTANT)
	// exists, err := h.accRepo.AccountingBookExists(
	// 	context.Background(),
	// 	event.OrganizationID,
	// )
	// if err != nil {
	// 	return err
	// }
	// if exists {
	// 	return nil // already created, safe
	// }

	// ✅ Create default accounting book
	book := &accounting_models.AccountingBook{
		ID:        uuid.New(),
		OrgID:     event.OrganizationID,
		Name:      "Default Accounting Book",
		CreatedAt: time.Now(),
	}

	return h.accRepo.CreateAccountingBook(
		ctx,
		book,
	)
}
