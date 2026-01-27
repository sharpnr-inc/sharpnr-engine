package accounting_service

type AccountingServiceInterface interface {

	// Setup Related stuffs
	CreateDraftAccountingSetup()

	SetupAccountingForLegalEntity()
}
