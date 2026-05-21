package legal_entity_repo

import (
	"context"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"

	"github.com/google/uuid"
)

func (d *orgLegalEntityRepo) AddLegalEntityIdentifier(ctx context.Context, req legal_entity_models.LegalEntityIdentifier) error {
	panic("")
}

func (d *orgLegalEntityRepo) FindAllIdentitifiersOfLegalEntity(ctx context.Context, legalEntityID uuid.UUID) ([]legal_entity_models.LegalEntityIdentifier, error) {
	panic("")
}

func (d *orgLegalEntityRepo) GetIdentifiersByID(ctx context.Context, identifierID uuid.UUID) (*legal_entity_models.LegalEntityIdentifier, error) {
	panic("")
}
