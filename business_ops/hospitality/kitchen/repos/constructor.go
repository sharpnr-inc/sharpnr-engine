package kitchenrepo

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func NewKitchenStagingRepository(stagingdb *gorm.DB) *kitchenStagingRepository {
	return &kitchenStagingRepository{
		db: stagingdb,
	}
}
