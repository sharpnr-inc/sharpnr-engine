package hotelservices

import (
	"context"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	hotelrepo "sharpnr-engine/modules/hospitality/hotel/hotel_repo"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/kitchenRepo"
	"sharpnr-engine/sere"
	"testing"

	"github.com/google/uuid"
)

func Test_hotelService_GetHotelDetails(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		sere        *sere.SERE
		hotelRepo   hotelrepo.HotelRepositoryInterface
		houseKeRepo hotelrepo.HotelHouseKeepingRepositoryInterface
		kitchenRepo kitchenrepo.KitchenRepositoryInterface
		// Named input parameters for target function.
		orgId   uuid.UUID
		want    *hospitalitydtos.GetHotel_RESPONSE
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewHotelService(tt.sere, tt.hotelRepo, tt.houseKeRepo, tt.kitchenRepo)
			got, gotErr := r.GetHotelDetails(context.Background(), tt.orgId)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetHotelDetails() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetHotelDetails() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetHotelDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
