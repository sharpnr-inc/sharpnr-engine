package legal_entity_service

import (
	"context"
	"fmt"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"

	"github.com/google/uuid"
)

// Replace the return types and arguments with whatever is actually
// defined in your OrgLegalEntityServiceInterface definition.
func (s *orgLegalEntity) GetTaxCategoryByID(ctx context.Context, taxCategoryID uuid.UUID) (legal_entity_contracts.TaxCategoryReader, error) {
	// Your implementation here...

	panic("implement me")
}

// Match the arguments and return values to your actual interface definition
func (s *orgLegalEntity) GetTaxCategoryDetails(ctx context.Context, taxCategoryID uuid.UUID) (legal_entity_contracts.TaxCategoryReader, error) {
	// TODO: Implement the actual repository call/logic here

	return legal_entity_contracts.TaxCategoryReader{}, fmt.Errorf("unimplemented: GetTaxCategoryDetails")
}
