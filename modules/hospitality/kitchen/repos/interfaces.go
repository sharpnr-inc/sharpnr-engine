package kitchenrepo

import (
	"context"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"

	"github.com/google/uuid"
)

type KitchenRepositoryInterface interface {
	CreateKitchen(ctx context.Context, kitchen *kitchen_models.Kitchen) error
	GetKitchenById(ctx context.Context, kitchenId uuid.UUID) ([]*kitchen_models.Kitchen, error)
	GetKitchenByOrgId(ctx context.Context, orgId uuid.UUID) ([]*kitchen_models.Kitchen, error)
	GetKitchenByPropertyId(ctx context.Context, propertyID uuid.UUID) ([]*kitchen_models.Kitchen, error)
	SwitchKitchenOnOff(ctx context.Context, KitchenID uuid.UUID, b bool) error

	// Ticket (KOT) Operations
	CreateKitchenTicket(ctx context.Context, ticket *kitchen_models.KitchenTicket) error
	CreateKitchenTicketItem(ctx context.Context, item *kitchen_models.KitchenTicketItem) error
	CreateKitchenTicketItemAddon(ctx context.Context, addon *kitchen_models.KitchenTicketItemAddon) error

	GetKitchenTicketsByKitchenID(ctx context.Context, kitchenID uuid.UUID) ([]*kitchen_models.KitchenTicket, error)
	GetKitchenTicketItemsByTicketID(ctx context.Context, kitchenTicketItemID uuid.UUID) ([]*kitchen_models.KitchenTicketItem, error)
	GetKitchenTicketItemAddOn(ctx context.Context, kitchenTicketItemID uuid.UUID) ([]*kitchen_models.KitchenTicketItemAddon, error)
}
