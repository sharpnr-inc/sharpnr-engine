package kitchenrepo

import (
	"context"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"

	"github.com/google/uuid"
)

type KitchenRepositoryInterface interface {
	CreateKitchen(ctx context.Context, kitchen *kitchen_models.Kitchen) error
	UpdateKitchenGeneralDetails(ctx context.Context, kitchenID uuid.UUID, updates map[string]interface{}) error
	GetKitchenById(ctx context.Context, kitchenId uuid.UUID) (*kitchen_models.Kitchen, error)
	GetKitchenByOrgId(ctx context.Context, orgId uuid.UUID) ([]*kitchen_models.Kitchen, error)
	GetKitchenByPropertyId(ctx context.Context, propertyID uuid.UUID) ([]*kitchen_models.Kitchen, error)
	SwitchKitchenOnOff(ctx context.Context, KitchenID uuid.UUID, b bool) error

	// Ticket (KOT) Operations
	CreateKitchenTicket(ctx context.Context, ticket *kitchen_models.KitchenTicket) error
	CreateKitchenTicketItem(ctx context.Context, item *kitchen_models.KitchenTicketItem) error
	CreateKitchenTicketItemVariant(ctx context.Context, item *kitchen_models.KitchenTicketItemVariant) error
	CreateKitchenTicketItemAddon(ctx context.Context, addon *kitchen_models.KitchenTicketItemAddon) error
	GetKOTByID(ctx context.Context, kotID uuid.UUID) (*kitchen_models.KitchenTicket, error)
	UpdateKotStatus(ctx context.Context, kotID uuid.UUID, status string) error

	GetKitchenTicketsByKitchenID(ctx context.Context, kitchenID uuid.UUID) ([]*kitchen_models.KitchenTicket, error)
	GetKitchenTicketItemsByTicketID(ctx context.Context, kitchenTicketItemID uuid.UUID) ([]*kitchen_models.KitchenTicketItem, error)
	GetKitchenTicketItemAddOn(ctx context.Context, kitchenTicketItemID uuid.UUID) ([]*kitchen_models.KitchenTicketItemAddon, error)
	GetKitchenTicketItemVariant(ctx context.Context, kitchenTicketItemID uuid.UUID) (*kitchen_models.KitchenTicketItemVariant, error)

	KitchenSettingsRepositoryInterface
}

type KitchenStagingRepositoryInterface interface {

	// =====================================================
	// KitchenTicketStaging CRUD
	// =====================================================

	CreateKitchenTicketStaging(
		ctx context.Context,
		ticket *kitchen_models.KitchenTicketStaging,
	) error

	GetKitchenTicketStagingByID(
		ctx context.Context,
		id uuid.UUID,
	) (*kitchen_models.KitchenTicketStaging, error)

	ListKitchenTicketStagingByKitchenID(
		ctx context.Context,
		kitchenID uuid.UUID,
	) ([]kitchen_models.KitchenTicketStaging, error)

	UpdateKitchenTicketStaging(
		ctx context.Context,
		ticket *kitchen_models.KitchenTicketStaging,
	) error

	DeleteKitchenTicketStaging(
		ctx context.Context,
		id uuid.UUID,
	) error

	// =====================================================
	// KitchenTicketItemStaging CRUD
	// =====================================================

	CreateKitchenTicketItemStaging(
		ctx context.Context,
		item *kitchen_models.KitchenTicketItemStaging,
	) error

	ListKitchenTicketItemStagingByTicketID(
		ctx context.Context,
		kitchenTicketID uuid.UUID,
	) ([]kitchen_models.KitchenTicketItemStaging, error)

	UpdateKitchenTicketItemStaging(
		ctx context.Context,
		item *kitchen_models.KitchenTicketItemStaging,
	) error

	DeleteKitchenTicketItemStaging(
		ctx context.Context,
		id uuid.UUID,
	) error

	// =====================================================
	// KitchenTicketItemVariantStaging CRUD
	// =====================================================

	CreateKitchenTicketItemVariantStaging(
		ctx context.Context,
		variant *kitchen_models.KitchenTicketItemVariantStaging,
	) error

	ListKitchenTicketItemVariantStagingByItemID(
		ctx context.Context,
		kitchenTicketItemID uuid.UUID,
	) ([]kitchen_models.KitchenTicketItemVariantStaging, error)

	DeleteKitchenTicketItemVariantStaging(
		ctx context.Context,
		id uuid.UUID,
	) error

	// =====================================================
	// KitchenTicketItemAddonStaging CRUD
	// =====================================================

	CreateKitchenTicketItemAddonStaging(
		ctx context.Context,
		addon *kitchen_models.KitchenTicketItemAddonStaging,
	) error

	ListKitchenTicketItemAddonStagingByItemID(
		ctx context.Context,
		kitchenTicketItemID uuid.UUID,
	) ([]kitchen_models.KitchenTicketItemAddonStaging, error)

	DeleteKitchenTicketItemAddonStaging(
		ctx context.Context,
		id uuid.UUID,
	) error
}

type KitchenSettingsRepositoryInterface interface {
	CreateKitchenAutoAcceptOrders(ctx context.Context, config *kitchen_models.KitchenAutoAcceptOrders) error
	CreateKitchenPrinterConfig(ctx context.Context, config *kitchen_models.KitchenPrinterConfig) error
	CreateKitchenThrottlingConfig(ctx context.Context, config *kitchen_models.KitchenThrottlingConfig) error
	CreateKitchenOperatingHours(ctx context.Context, config *kitchen_models.KitchenOperatingHours) error

	GetKitchenAutoAcceptOrdersByKitchenID(ctx context.Context, kitchenID uuid.UUID) (*kitchen_models.KitchenAutoAcceptOrders, error)
	GetKitchenPrinterConfigByKitchenID(ctx context.Context, kitchenID uuid.UUID) (*kitchen_models.KitchenPrinterConfig, error)
	GetKitchenThrottlingConfigByKitchenID(ctx context.Context, kitchenID uuid.UUID) (*kitchen_models.KitchenThrottlingConfig, error)
	ListKitchenOperatingHoursByKitchenID(ctx context.Context, kitchenID uuid.UUID) ([]kitchen_models.KitchenOperatingHours, error)

	UpdateKitchenAutoAcceptOrders(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	UpdateKitchenPrinterConfig(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	UpdateKitchenThrottlingConfig(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	UpdateKitchenOperatingHours(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
}
