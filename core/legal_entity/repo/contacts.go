package legal_entity_repo

import (
	"context"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"

	"github.com/google/uuid"
)

func (d *orgLegalEntityRepo) AddLegalEntityContacts(ctx context.Context, req legal_entity_models.LegalEntityContact) error {
	panic("")
}

func (d *orgLegalEntityRepo) FindAddAllContactsOfLegalEntity(ctx context.Context, legalEntityID uuid.UUID) ([]legal_entity_models.LegalEntityContact, error) {
	panic("")
}
