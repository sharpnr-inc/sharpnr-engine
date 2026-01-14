package kitchenrepo

import (
	"context"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"

	"github.com/google/uuid"
)

type KitchenRepositoryInterface interface {
	//This function for Creating the Create Kitchen
	// this function used in the Org factory to create default Kitchen
	CreateKitchen(ctx context.Context, kitchen *kitchen_models.Kitchen) error

	GetKitchenByOrgId(ctx context.Context, orgId uuid.UUID) (*kitchen_models.Kitchen, error)

	// Kitchen Tickets
	CreateKitchenTicketDBTx(ctx context.Context, ticket *kitchen_models.KitchenTicket) (*kitchen_models.KitchenTicket, error)
	GetKitchenTicketByIdDBTx(ctx context.Context, ticketId uuid.UUID) (*kitchen_models.KitchenTicket, error)

	CreateKitchenTicketItemDBTx(ctx context.Context, item *kitchen_models.KitchenTicketItem) (*kitchen_models.KitchenTicketItem, error)

	CreateKitchenTicketItemAddonDBTx(
		ctx context.Context,
		addon *kitchen_models.KitchenTicketItemAddon,
	) (*kitchen_models.KitchenTicketItemAddon, error)

	GetNextTicketNumber(
		ctx context.Context,
		kitchenId uuid.UUID,
	) (string, error)

	GetKitchenTickets(ctx context.Context, kitchenId uuid.UUID) ([]kitchen_models.KitchenTicket, error)
}
