package kitchenrepo

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	kitchen_models "sharpnr-engine/modules/hospitality/kitchen/models"
)

type kitchenRepository struct {
	db *gorm.DB
}

func NewKitchenRepository(db *gorm.DB) *kitchenRepository {
	return &kitchenRepository{
		db: db,
	}
}

type kitchenStagingRepository struct {
	db *gorm.DB
}

// CreateKitchenTicketItemAddonStaging implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) CreateKitchenTicketItemAddonStaging(ctx context.Context, addon *kitchen_models.KitchenTicketItemAddonStaging) error {
	panic("unimplemented")
}

// CreateKitchenTicketItemStaging implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) CreateKitchenTicketItemStaging(ctx context.Context, item *kitchen_models.KitchenTicketItemStaging) error {
	panic("unimplemented")
}

// CreateKitchenTicketItemVariantStaging implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) CreateKitchenTicketItemVariantStaging(ctx context.Context, variant *kitchen_models.KitchenTicketItemVariantStaging) error {
	panic("unimplemented")
}

// CreateKitchenTicketStaging implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) CreateKitchenTicketStaging(ctx context.Context, ticket *kitchen_models.KitchenTicketStaging) error {
	panic("unimplemented")
}

// DeleteKitchenTicketItemAddonStaging implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) DeleteKitchenTicketItemAddonStaging(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteKitchenTicketItemStaging implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) DeleteKitchenTicketItemStaging(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteKitchenTicketItemVariantStaging implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) DeleteKitchenTicketItemVariantStaging(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// DeleteKitchenTicketStaging implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) DeleteKitchenTicketStaging(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetKitchenTicketStagingByID implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) GetKitchenTicketStagingByID(ctx context.Context, id uuid.UUID) (*kitchen_models.KitchenTicketStaging, error) {
	panic("unimplemented")
}

// ListKitchenTicketItemAddonStagingByItemID implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) ListKitchenTicketItemAddonStagingByItemID(ctx context.Context, kitchenTicketItemID uuid.UUID) ([]kitchen_models.KitchenTicketItemAddonStaging, error) {
	panic("unimplemented")
}

// ListKitchenTicketItemStagingByTicketID implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) ListKitchenTicketItemStagingByTicketID(ctx context.Context, kitchenTicketID uuid.UUID) ([]kitchen_models.KitchenTicketItemStaging, error) {
	panic("unimplemented")
}

// ListKitchenTicketItemVariantStagingByItemID implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) ListKitchenTicketItemVariantStagingByItemID(ctx context.Context, kitchenTicketItemID uuid.UUID) ([]kitchen_models.KitchenTicketItemVariantStaging, error) {
	panic("unimplemented")
}

// ListKitchenTicketStagingByKitchenID implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) ListKitchenTicketStagingByKitchenID(ctx context.Context, kitchenID uuid.UUID) ([]kitchen_models.KitchenTicketStaging, error) {
	panic("unimplemented")
}

// UpdateKitchenTicketItemStaging implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) UpdateKitchenTicketItemStaging(ctx context.Context, item *kitchen_models.KitchenTicketItemStaging) error {
	panic("unimplemented")
}

// UpdateKitchenTicketStaging implements [KitchenStagingRepositoryInterface].
func (k *kitchenStagingRepository) UpdateKitchenTicketStaging(ctx context.Context, ticket *kitchen_models.KitchenTicketStaging) error {
	panic("unimplemented")
}

func NewKitchenStagingRepository(stagingdb *gorm.DB) *kitchenStagingRepository {
	return &kitchenStagingRepository{
		db: stagingdb,
	}
}
