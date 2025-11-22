package hotelrepo

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Test_hotelRepository_UpdateRoom(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		db *gorm.DB
		// Named input parameters for target function.
		fields  map[string]interface{}
		roomId  uuid.UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHotelRepository(tt.db)
			gotErr := h.UpdateRoom(context.Background(), tt.fields, tt.roomId)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateRoom() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateRoom() succeeded unexpectedly")
			}
		})
	}
}
