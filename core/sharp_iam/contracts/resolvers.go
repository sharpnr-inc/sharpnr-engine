package contracts

import (
	"context"
	"github.com/google/uuid"
)

type ParentNode struct {
	Type string
	ID   uuid.UUID
}

type ResourceHierarchyResolver interface {
	ResolveParents(
		ctx context.Context,
		orgID uuid.UUID,
		resourceType string,
		resourceID uuid.UUID,
	) ([]ParentNode, error)

	ResolveAttributes(
		ctx context.Context,
		orgID uuid.UUID,
		resourceType string,
		resourceID uuid.UUID,
	) (map[string]string, error)
}
