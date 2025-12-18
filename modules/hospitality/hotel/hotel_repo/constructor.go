package hotelrepo

import (
	"context"
	hospitalitymodels "sharpnr-engine/modules/hospitality/hospitality_models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hotelRepository struct {
	db *gorm.DB
}

type hotelHouseKeepingRepository struct {
	db *gorm.DB
}

// GetAllRoomInFloor implements HotelRepository.
func (h *hotelRepository) GetAllRoomInFloor() {
	panic("unimplemented")
}

// GetFloorsByHotel implements HotelRepository.
func (h *hotelRepository) GetFloorsByHotel(ctx context.Context, hotelID uuid.UUID) ([]hospitalitymodels.HotelFloor, error) {
	panic("unimplemented")
}

// UpdateFloor implements HotelRepository.
func (h *hotelRepository) UpdateFloor(ctx context.Context, floor *hospitalitymodels.HotelFloor) error {
	panic("unimplemented")
}

func (r *hotelRepository) UpdateHotelFields(
	ctx context.Context,
	hotelID uuid.UUID,
	updates map[string]interface{},
) error {

	if len(updates) == 0 {
		return nil
	}

	return r.db.
		WithContext(ctx).
		Model(&hospitalitymodels.Hotel{}).
		Where("id = ?", hotelID).
		Updates(updates).
		Error
}

//-----------------------------------------------------------------------------------------------------------------------------------//

// NewHotelRepository creates a new instance of hotelRepository.
//
// It initializes the repository with a gorm.DB connection and returns
// an implementation of the HotelRepository interface.
//
// This constructor should be used by the service layer or dependency
// injection setup to ensure a single, consistent database connection
// is used across all hotel-related persistence operations.
func NewHotelRepository(db *gorm.DB) *hotelRepository {
	return &hotelRepository{
		db: db,
	}
}

// WithTx runs multiple repository operations inside a single database transaction.
//
// What this does:
//   - Starts a database transaction
//   - Provides a transaction-scoped repository to the caller
//   - Commits the transaction if the function returns nil
//   - Rolls back the transaction if the function returns an error
//
// Why this is needed:
//
//	Some business operations require inserting or updating multiple tables
//	together (for example: booking + guest + payment).
//	Using a transaction guarantees that either ALL changes are saved,
//	or NONE of them are saved.
//
// How to use:
//
//	err := repo.WithTx(ctx, func(txRepo HotelRepositoryTransactions) error {
//	    if err := txRepo.CreateRoom(ctx, room); err != nil {
//	        return err // triggers rollback
//	    }
//
//	    if err := txRepo.CreateRoomPrice(ctx, price); err != nil {
//	        return err // triggers rollback
//	    }
//
//	    return nil // commits the transaction
//	})
//
// Important notes:
//   - Do NOT start or commit transactions manually inside services
//   - All operations inside the callback MUST use txRepo
//   - Returning an error automatically rolls back the transaction
//   - Returning nil commits the transaction
func (r *hotelRepository) WithTx(
	ctx context.Context,
	fn func(txRepo *hotelRepository) error,
) error {

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create a new repository instance bound to the transaction
		txRepo := &hotelRepository{db: tx}
		return fn(txRepo)
	})
}

// -----------------------------------------------------------------------------------------------------------------------------------//

func NewHotelHouseKeepingRepository(db *gorm.DB) *hotelHouseKeepingRepository {
	return &hotelHouseKeepingRepository{
		db: db,
	}
}
