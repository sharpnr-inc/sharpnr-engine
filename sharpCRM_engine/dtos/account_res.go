package sharpCrmEngineDtos

import "github.com/google/uuid"

type SharpCrmAccountStatusResponse struct {
	ServiceStarted bool
	IsActive       *bool
	ID             *uuid.UUID
}
