package kitchenrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
)

type KitchenRepositoryInterface interface {
	//This function for Creating the Create Kitchen
	// this function used in the Org factory to create default Kitchen
	CreateKitchen(ctx context.Context, kitchen *hospitalitymodels.Kitchen) error

	GetKitchenByOrgId(ctx context.Context, orgId uuid.UUID) (*hospitalitymodels.Kitchen, error)

	// Kitchen Tickets
	CreateKitchenTicketDBTx(ctx context.Context, ticket *hospitalitymodels.KitchenTicket) (*hospitalitymodels.KitchenTicket, error)
	GetKitchenTicketByIdDBTx(ctx context.Context, ticketId uuid.UUID) (*hospitalitymodels.KitchenTicket, error)

	CreateKitchenTicketItemDBTx(ctx context.Context, item *hospitalitymodels.KitchenTicketItem) (*hospitalitymodels.KitchenTicketItem, error)

	CreateKitchenTicketItemAddonDBTx(
		ctx context.Context,
		addon *hospitalitymodels.KitchenTicketItemAddon,
	) (*hospitalitymodels.KitchenTicketItemAddon, error)

	GetNextTicketNumber(
		ctx context.Context,
		kitchenId uuid.UUID,
	) (string, error)

	GetKitchenTickets(ctx context.Context, kitchenId uuid.UUID) ([]hospitalitymodels.KitchenTicket, error)
}
