package iamEngineService

import (
	"context"

	iamEngineDtos "sharpnr-engine/iam_engine/dtos"

	"github.com/google/uuid"
)

type IamEngineServiceInterface interface {
	GetProfile(ctx context.Context, orgID uuid.UUID, empID uuid.UUID) (*iamEngineDtos.ProfileResponse_dto, error)
}
