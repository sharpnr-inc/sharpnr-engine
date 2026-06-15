package legal_entity_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	legal_entity_contracts "sharpnr-engine/core/legal_entity/contracts"
	legal_entity_models "sharpnr-engine/core/legal_entity/models"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/shared/utils"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// Create a new Legal Entity for the organization
// and returns the LegalEntityID
func (r *orgLegalEntity) CreateNewLegalEntity(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	input legal_entity_contracts.NewLegalEntityRequestDto,
) (legalEntityID uuid.UUID, err error) {

	err = r.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		incorporationDate, err := utils.ConvertDate(input.IncorporationDate)
		if err != nil {
			return err
		}

		legalEntity := &legal_entity_models.LegalEntity{
			ID:                uuid.New(),
			OrganizationID:    orgID,
			LegalName:         input.LegalName,
			DisplayName:       input.DisplayName,
			EntityType:        input.Type,
			IsActive:          false,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
			FYStartMonth:      input.FiscalYearMonth,
			IncorporationDate: &incorporationDate,
		}

		if err := r.legalEntityRepo.CreateLegalEntity(txCtx, legalEntity); err != nil {
			return err
		}

		if len(input.Addresses) > 0 {
			for _, a := range input.Addresses {
				_, err = r.AddNewAddressInLegalEntity(txCtx, orgID, emplID, legalEntity.ID, a)
				if err != nil {
					return err
				}
			}
		}

		legalEntityID = legalEntity.ID
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	return legalEntityID, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orgLegalEntity) GetLegalEntityDetails(
	ctx context.Context,
	orgID, emplId, legalEntityID uuid.UUID,
) (legal_entity_contracts.LegalEntityDetails, error) {

	var result legal_entity_contracts.LegalEntityDetails

	// 1. Fetch Core Details sequentially first (Acts as an existential validation guard)
	le, err := s.legalEntityRepo.GetLegalEntityByID(ctx, legalEntityID)
	if err != nil {
		return result, err
	}
	if le == nil {
		return result, fmt.Errorf("legal entity not found: %s", legalEntityID)
	}

	// Map basic details immediately
	result.BasicDetails = legal_entity_contracts.LegalEntityResponseDto{
		ID:                le.ID,
		DisplayName:       le.DisplayName,
		LegalName:         le.LegalName,
		Status:            le.IsActive,
		EntityType:        le.EntityType,
		IncorporationDate: *le.IncorporationDate,
	}

	// 2. Initialize the error group using the request context thread safety
	g, groupCtx := errgroup.WithContext(ctx)

	// Intermediate memory locations to hold concurrent workers data safely
	var (
		leContacts        []legal_entity_models.LegalEntityContact
		leBankAccounts    []legal_entity_models.LegalEntityBankAccount
		leAddr            []legal_entity_models.LegalEntityAddress
		leLicenseComp     []*legal_entity_models.LegalEntityCompliance
		leIdentifiers     []legal_entity_models.LegalEntityIdentifier
		leTaxRegistration []legal_entity_models.LegalEntityTaxRegistration
		taxCategory       []legal_entity_contracts.LegalEntityTaxCategoryResponseDto
	)

	// Parallel Worker 1: Contacts
	g.Go(func() error {
		var err error
		leContacts, err = s.legalEntityRepo.FindAddAllContactsOfLegalEntity(groupCtx, legalEntityID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil // Soft fallback: Keep group context alive, array remains empty slice
		}
		return err
	})

	// Parallel Worker 2: Bank Accounts
	g.Go(func() error {
		var err error
		leBankAccounts, err = s.legalEntityRepo.FindBankAccountsByLegalEntityID(groupCtx, legalEntityID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	})

	// Parallel Worker 3: Addresses
	g.Go(func() error {
		var err error
		leAddr, err = s.legalEntityRepo.GetAddressesOfLegalEntity(groupCtx, le.ID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	})

	// Parallel Worker 4: Compliances & Licenses
	g.Go(func() error {
		var err error
		leLicenseComp, err = s.legalEntityRepo.FindAllLegalCompliancesByLegalEntityID(groupCtx, legalEntityID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	})

	// Parallel Worker 5: Identifiers
	g.Go(func() error {
		var err error
		leIdentifiers, err = s.legalEntityRepo.FindAllIdentitifiersOfLegalEntity(groupCtx, legalEntityID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	})

	// Parallel Worker 6: Tax Registrations
	g.Go(func() error {
		var err error
		leTaxRegistration, err = s.legalEntityRepo.FindAllTaxRegistrationOfLegalEntity(groupCtx, legalEntityID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	})

	// Parallel Worker 7: Nested Tax Categories
	g.Go(func() error {
		var err error
		taxCategory, err = s.GetLegalEntityTaxCategory(groupCtx, orgID, emplId, legalEntityID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	})

	// 3. Wait block: Block routine until workers yield or any genuine critical DB error hits
	if err := g.Wait(); err != nil {
		return legal_entity_contracts.LegalEntityDetails{}, err
	}

	// 4. Initialize result arrays to empty slices to safeguard frontend rendering against null payloads
	result.Contacts = []legal_entity_contracts.LegalEntityContactResponseDto{}
	result.BankAccounts = []legal_entity_contracts.LegalEntityBankAccountsResponseDto{}
	result.AddressesAndOffices = []legal_entity_contracts.LegalEntityAddressResponseDto{}
	result.LicencesAndCompliance = []legal_entity_contracts.LegalEntityComplianceResponseDto{}
	result.Identifiers = []legal_entity_contracts.LegalEntityIdentifierResponseDto{}
	result.TaxRegistrations = []legal_entity_contracts.LegalEntityTaxRegistrationResponseDto{}
	result.TaxCategories = taxCategory // Already initialized to [] internally inside GetLegalEntityTaxCategory

	// 5. Structure and Map database models over to Response DTO formats
	for _, c := range leContacts {
		result.Contacts = append(result.Contacts, legal_entity_contracts.LegalEntityContactResponseDto{
			ID:          c.ID,
			Type:        c.Type,
			FirstName:   c.FirstName,
			LastName:    c.LastName,
			Designation: c.Designation,
			Email:       c.Email,
			Mobile:      c.Mobile,
			Din:         c.DIN,
			IsPrimary:   c.IsPrimary,
			IsActive:    c.IsActive,
		})
	}

	for _, b := range leBankAccounts {
		result.BankAccounts = append(result.BankAccounts, legal_entity_contracts.LegalEntityBankAccountsResponseDto{
			ID:                b.ID,
			AccountHolderName: b.AccountHolderName,
			AccountNumber:     b.AccountNumber,
			IFSC:              b.IFSC,
			AccountType:       b.AccountType,
			IsPrimary:         b.IsPrimary,
			IsVerified:        b.IsVerified,
			VerifiedAt:        b.VerifiedAt,
			IsActive:          b.IsActive,
			AddedAt:           b.CreatedAt,
		})
	}

	for _, a := range leAddr {
		result.AddressesAndOffices = append(result.AddressesAndOffices, legal_entity_contracts.LegalEntityAddressResponseDto{
			ID:           a.ID,
			Type:         a.Type,
			AddressLine1: a.AddressLine1,
			AddressLine2: a.AddressLine2,
			Landmark:     a.Landmark,
			City:         a.City,
			State:        a.State,
			GSTStateCode: a.GSTStateCode,
			Pincode:      a.Pincode,
			Country:      a.Country,
			IsPrimary:    a.IsPrimary,
			PropertyID:   a.PropertyID,
		})
	}

	for _, lc := range leLicenseComp {
		if lc == nil {
			continue
		}
		result.LicencesAndCompliance = append(result.LicencesAndCompliance, legal_entity_contracts.LegalEntityComplianceResponseDto{
			ID:                  lc.ID,
			Type:                lc.Type,
			LicenseNumber:       lc.LicenseNumber,
			IssuingAuthority:    lc.IssuingAuthority,
			ValidFrom:           lc.ValidFrom,
			ValidTo:             lc.ValidTo,
			RenewalReminderDays: lc.RenewalReminderDays,
			PropertyID:          lc.PropertyID,
			Module:              lc.Module,
			DocumentURL:         lc.DocumentURL,
			IsActive:            lc.IsActive,
		})
	}

	for _, id := range leIdentifiers {
		result.Identifiers = append(result.Identifiers, legal_entity_contracts.LegalEntityIdentifierResponseDto{
			ID:         id.ID,
			Type:       id.Type,
			Value:      id.Value,
			IsVerified: id.IsVerified,
			VerifiedAt: id.VerifiedAt,
			ExpiresAt:  id.ExpiresAt,
		})
	}

	for _, tr := range leTaxRegistration {
		result.TaxRegistrations = append(result.TaxRegistrations, legal_entity_contracts.LegalEntityTaxRegistrationResponseDto{
			ID:                  tr.ID,
			TaxType:             tr.TaxType,
			RegistrationNumber:  tr.RegistrationNumber,
			GSTRegistrationType: tr.GSTRegistrationType,
			State:               tr.State,
			GSTStateCode:        tr.GSTStateCode,
			Country:             tr.Country,
			EffectiveFrom:       tr.EffectiveFrom,
			EffectiveTo:         tr.EffectiveTo,
			IsActive:            tr.IsActive,
		})
	}

	return result, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orgLegalEntity) GetLegalEntityByID(ctx context.Context, orgID, emplID, legalEntityID uuid.UUID) (*legal_entity_contracts.LegalEntityResponseDto, error) {

	legalEntity, err := s.legalEntityRepo.GetLegalEntityByID(ctx, legalEntityID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	return &legal_entity_contracts.LegalEntityResponseDto{
		ID:          legalEntity.ID,
		DisplayName: legalEntity.DisplayName,
	}, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orgLegalEntity) GetLegalEntityByOrgID(
	ctx context.Context,
	orgID uuid.UUID,
	emplID uuid.UUID,
) ([]legal_entity_contracts.LegalEntityResponseDto, error) {

	var results []legal_entity_contracts.LegalEntityResponseDto

	err := s.txManager.WithTx(ctx, func(tx *gorm.DB) error {

		txCtx := dbtx.InjectTx(ctx, tx)

		// Fetch legal entities
		legalEntities, err := s.legalEntityRepo.GetLegalEntitiesByOrgID(txCtx, orgID)
		if err != nil {
			return err
		}

		results = make([]legal_entity_contracts.LegalEntityResponseDto, len(legalEntities))

		for i, entity := range legalEntities {
			var incDate time.Time
			if entity.IncorporationDate != nil {
				incDate = *entity.IncorporationDate
			}

			results[i] = legal_entity_contracts.LegalEntityResponseDto{
				ID:                entity.ID,
				DisplayName:       entity.DisplayName,
				LegalName:         entity.LegalName,
				Status:            entity.IsActive, // Maps IsActive to Status
				EntityType:        entity.EntityType,
				IncorporationDate: incDate,
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return results, nil
}

/* ---------------------------------- <***> --------------------------------- */

func (s *orgLegalEntity) GetLegalEntityTaxCategory(
	ctx context.Context,
	orgID, emplID, legalEntityID uuid.UUID,
) ([]legal_entity_contracts.LegalEntityTaxCategoryResponseDto, error) {

	// Safeguard: Initialize as an empty slice so frontend gets [] instead of null if empty
	results := []legal_entity_contracts.LegalEntityTaxCategoryResponseDto{}

	// 1. Fetch the root Tax Categories for this Legal Entity
	leTaxCategories, err := s.legalEntityRepo.ListTaxCategories(ctx, legalEntityID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return results, nil
		}
		return nil, fmt.Errorf("failed to list tax categories: %w", err)
	}

	// 2. Build the nested tree hierarchy sequentially
	for _, tc := range leTaxCategories {
		categoryDto := legal_entity_contracts.LegalEntityTaxCategoryResponseDto{
			ID:              tc.ID,
			Name:            tc.Name,
			Code:            tc.Code,
			Description:     tc.Description,
			SACCode:         tc.SACCode,
			HSNCode:         tc.HSNCode,
			IsActive:        tc.IsActive,
			Rules:           []legal_entity_contracts.LegalEntityTaxRuleResponseDto{},
			AccountMappings: []legal_entity_contracts.TaxCategoryAccountMappingResponseDto{},
		}

		// A. Fetch and process rules for this tax category
		taxCategoryRules, err := s.legalEntityRepo.FindAllTaxCategoryRules(ctx, tc.ID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("failed to fetch rules for tax category %s: %w", tc.ID, err)
		}

		for _, trc := range taxCategoryRules {
			ruleDto := legal_entity_contracts.LegalEntityTaxRuleResponseDto{
				ID:                trc.ID,
				TaxRegistrationID: trc.LegalEntityTaxRegistrationID,
				TaxType:           trc.TaxType,
				Rate:              trc.Rate,
				ConditionJSON:     json.RawMessage(trc.ConditionJSON), // Direct cast to json.RawMessage
				MinValue:          trc.MinValue,
				MaxValue:          trc.MaxValue,
				Priority:          trc.Priority,
				EffectiveFrom:     trc.EffectiveFrom,
				EffectiveTo:       trc.EffectiveTo,
				IsActive:          trc.IsActive,
				Components:        []legal_entity_contracts.LegalEntityTaxComponentResponseDto{},
			}

			// B. Fetch components for this specific rule
			taxRuleComponents, err := s.legalEntityRepo.GetTaxRuleComponents(ctx, trc.ID)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("failed to fetch components for rule %s: %w", trc.ID, err)
			}

			for _, comp := range taxRuleComponents {
				compDto := legal_entity_contracts.LegalEntityTaxComponentResponseDto{
					ID:                 comp.ID,
					ComponentCode:      comp.ComponentCode,
					Rate:               comp.Rate,
					LiabilityAccountID: comp.LiabilityAccountID,
				}
				ruleDto.Components = append(ruleDto.Components, compDto)
			}

			categoryDto.Rules = append(categoryDto.Rules, ruleDto)
		}

		// C. Fetch Account Mapping for this tax category
		accMap, err := s.legalEntityRepo.GetTaxCategoryAccountMappingByTaxCategoryID(ctx, tc.ID)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("failed to fetch account mapping for tax category %s: %w", tc.ID, err)
			}
			// If mapping row doesn't exist, it is perfectly fine — skip mapping addition
		} else if accMap != nil {
			mappingDto := legal_entity_contracts.TaxCategoryAccountMappingResponseDto{
				ID:                  accMap.ID,
				AccountingBookID:    accMap.AccountingBookID,
				ReceivableAccountID: accMap.ReceivableAccountID,
				IncomeAccountID:     accMap.IncomeAccountID,
				IsActive:            accMap.IsActive,
			}
			categoryDto.AccountMappings = append(categoryDto.AccountMappings, mappingDto)
		}

		results = append(results, categoryDto)
	}

	return results, nil
}
