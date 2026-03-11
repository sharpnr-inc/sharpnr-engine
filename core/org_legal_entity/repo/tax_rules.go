package legal_entity_repo

import (
	"context"
	dbtx "sharpnr-engine/core/dbTx"
	legal_entity_models "sharpnr-engine/core/org_legal_entity/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *orgLegalEntityRepo) CreateTaxRuleWithComponents(
	ctx context.Context,
	rule *legal_entity_models.LegalEntityTaxRule,
	components []legal_entity_models.LegalEntityTaxComponent,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Wrap in a transaction to ensure both Rule and Components are saved together
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if rule.ID == uuid.Nil {
			rule.ID = uuid.New()
		}

		if err := tx.Create(rule).Error; err != nil {
			return err
		}

		// Link components to the newly created Rule
		for i := range components {
			if components[i].ID == uuid.Nil {
				components[i].ID = uuid.New()
			}
			components[i].LegalEntityTaxRuleID = rule.ID
		}

		if len(components) > 0 {
			if err := tx.Create(&components).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *orgLegalEntityRepo) GetActiveRulesByCategory(ctx context.Context, categoryID uuid.UUID) ([]legal_entity_models.LegalEntityTaxRule, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var rules []legal_entity_models.LegalEntityTaxRule
	now := time.Now()

	err := db.WithContext(ctx).
		Where("legal_entity_tax_category_id = ?", categoryID).
		Where("is_active = ?", true).
		// Check for validity period: Start date passed and (End date is null OR End date hasn't passed)
		Where("effective_from <= ?", now).
		Where("(effective_to IS NULL OR effective_to >= ?)", now).
		Order("effective_from DESC").
		Find(&rules).
		Error

	return rules, err
}

func (r *orgLegalEntityRepo) GetRuleComponents(ctx context.Context, ruleID uuid.UUID) ([]legal_entity_models.LegalEntityTaxComponent, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var components []legal_entity_models.LegalEntityTaxComponent

	err := db.WithContext(ctx).
		Where("legal_entity_tax_rule_id = ?", ruleID).
		Order("component_code ASC").
		Find(&components).
		Error

	return components, err
}

// --- Accounting Mapping (The "Where" in the Ledger) ---
// Links a Tax Category to specific Ledger Accounts in a specific Book

func (r *orgLegalEntityRepo) UpsertAccountingMapping(ctx context.Context, mapping *legal_entity_models.LegalEntityAccountingMapping) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// .Save handles both INSERT and UPDATE based on the Primary Key (ID)
	return db.WithContext(ctx).Save(mapping).Error
}

func (r *orgLegalEntityRepo) GetMappingByCategory(ctx context.Context, categoryID uuid.UUID, bookID uuid.UUID) (*legal_entity_models.LegalEntityAccountingMapping, error) {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	var mapping legal_entity_models.LegalEntityAccountingMapping

	err := db.WithContext(ctx).
		Where("legal_entity_tax_category_id = ? AND accounting_book_id = ?", categoryID, bookID).
		Where("is_active = ?", true).
		First(&mapping).
		Error

	if err != nil {
		return nil, err
	}

	return &mapping, nil
}

func (r *orgLegalEntityRepo) ToggleRuleStatus(ctx context.Context, ruleID uuid.UUID, isActive bool) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// Using .Model and .Update for a partial update (efficient)
	return db.WithContext(ctx).
		Model(&legal_entity_models.LegalEntityTaxRule{}).
		Where("id = ?", ruleID).
		Update("is_active", isActive).
		Error
}
