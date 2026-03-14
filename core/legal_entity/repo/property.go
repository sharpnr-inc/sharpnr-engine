package legal_entity_repo

import (
	"context"
	"fmt"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
)

func (o *orgLegalEntityRepo) AddPropertyInLegalEntity(ctx context.Context, relation *legal_entity_models.PropertyLegalEntity) error {
	// 1. Set defaults for a new relationship
	relation.IsActive = true

	// 2. Insert the relationship row
	// If you have a Unique Index on (PropertyID, LegalEntityID, Role),
	// this will fail if the exact relationship already exists.
	err := o.db.WithContext(ctx).Create(relation).Error

	if err != nil {
		return fmt.Errorf("failed to link property to legal entity: %w", err)
	}

	return nil
}
