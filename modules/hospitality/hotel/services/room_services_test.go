package hotelservices

import (
	"context"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/sere"
	"testing"

	"github.com/google/uuid"
)

func Test_hotelService_CreateRoomInHotel(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		sere        *sere.SERE
		hotelRepo   hotelrepo.HotelRepositoryInterface
		houseKeRepo hotelrepo.HotelHouseKeepingRepositoryInterface
		kitchenRepo kitchenrepo.KitchenRepositoryInterface
		// Named input parameters for target function.
		orgID   uuid.UUID
		req     hoteldtos.CreateRoomInFloor_REQ
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHotelService(tt.sere, tt.hotelRepo, tt.houseKeRepo, tt.kitchenRepo)
			gotErr := h.CreateRoomInHotel(context.Background(), tt.orgID, tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateRoomInHotel() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateRoomInHotel() succeeded unexpectedly")
			}
		})
	}
}
