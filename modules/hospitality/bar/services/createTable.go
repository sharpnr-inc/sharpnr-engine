package barservices

import (
	"context"
	"errors"
	"fmt"
	bar_models "sharpnr-engine/modules/hospitality/bar/models"
	hospitalitydtos "sharpnr-engine/modules/hospitality/hospitality_dtos"
	"time"

	"github.com/google/uuid"
)

func (s *barServices) CreateTablesInBar(ctx context.Context, orgId uuid.UUID, req hospitalitydtos.Bar_CreateTablesInBar_Request) error {

	bar, err := s.barRepo.GetBarByOrgId(ctx, orgId)
	if err != nil {
		return err
	}
	if bar.OrganizationID != orgId {
		return errors.New("bar not found")
	}

	// 2️⃣ Validate Floor belongs to this Bar
	floor, err := s.barRepo.GetFloorById(ctx, req.FloorId)
	if err != nil {
		return fmt.Errorf("floor not found: %w", err)
	}

	if floor.BarID != req.BarId {
		return errors.New("floor does not belong to the specified bar")
	}

	// 3️⃣ Create tables inside this floor
	for _, tbl := range req.TableList {

		table := &bar_models.BarTable{
			ID:        uuid.New(),
			FloorID:   req.FloorId,
			Floor:     floor,
			Name:      tbl.TableName,
			Capacity:  tbl.TableCapacity,
			Status:    "AVAILABLE",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := s.barRepo.CreateTable(ctx, table); err != nil {
			return fmt.Errorf("failed to create table '%s': %w", tbl.TableName, err)
		}
	}

	return nil
}
