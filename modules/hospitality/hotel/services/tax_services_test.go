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

func Test_hotelService_CreateTaxRule(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		sere        *sere.SERE
		hotelRepo   hotelrepo.HotelRepositoryInterface
		houseKeRepo hotelrepo.HotelHouseKeepingRepositoryInterface
		kitchenRepo kitchenrepo.KitchenRepositoryInterface
		// Named input parameters for target function.
		orgId   uuid.UUID
		req     hoteldtos.CreateHotelTaxRuleRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewHotelService(tt.sere, tt.hotelRepo, tt.houseKeRepo, tt.kitchenRepo)
			gotErr := s.CreateTaxRule(context.Background(), tt.orgId, tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateTaxRule() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateTaxRule() succeeded unexpectedly")
			}
		})
	}
}
