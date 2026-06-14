package legal_entity_repo

import (
	"context"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *orgLegalEntityRepo) CreateTaxRuleOfTaxCategory(
	ctx context.Context,
	rule *legal_entity_models.LegalEntityTaxRule,
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

		return nil
	})
}

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) CreateTaxComponentOfTaxRule(
	ctx context.Context,
	tax_components *legal_entity_models.LegalEntityTaxComponent,
) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	return db.WithContext(ctx).Create(tax_components).Error
}

/* ---------------------------------- <***> --------------------------------- */

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

/* ---------------------------------- <***> --------------------------------- */

func (r *orgLegalEntityRepo) GetTaxRuleComponents(ctx context.Context, ruleID uuid.UUID) ([]legal_entity_models.LegalEntityTaxComponent, error) {
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

/* ---------------------------------- <***> --------------------------------- */

// --- Accounting Mapping (The "Where" in the Ledger) ---
// Links a Tax Category to specific Ledger Accounts in a specific Book
func (r *orgLegalEntityRepo) UpsertAccountingMapping(ctx context.Context, mapping *legal_entity_models.TaxCategoryAccountMapping) error {
	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = r.db
	}

	// .Save handles both INSERT and UPDATE based on the Primary Key (ID)
	return db.WithContext(ctx).Save(mapping).Error
}

/* ---------------------------------- <***> --------------------------------- */

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

/* ---------------------------------- <***> --------------------------------- */

func (d *orgLegalEntityRepo) FindAllTaxCategoryRules(
	ctx context.Context,
	taxCategoryID uuid.UUID,
) ([]legal_entity_models.LegalEntityTaxRule, error) {

	db, ok := dbtx.ExtractTx(ctx)
	if !ok {
		db = d.db
	}

	var taxRules []legal_entity_models.LegalEntityTaxRule

	err := db.WithContext(ctx).
		Where("tax_category_id = ?", taxCategoryID).
		Find(&taxRules).
		Error

	if err != nil {
		return nil, err
	}

	// Explicitly return gorm.ErrRecordNotFound if no rules are found for this category
	if len(taxRules) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return taxRules, nil
}
