package accounting_service

// Accounting's read across into the stock engine. Parked: the method it called
// went with the inventory tables migration 000062 dropped, and the stockwave
// rewrite has not published a replacement yet. The dependency is still wired
// (accountingServices.stockwaveEngineService), so this comes back as a real
// method the day stockwave exposes the per-legal-entity stock read again.
//
// func (s *accountingServices) GetAllInventoriesAccounts(ctx context.Context, orgID, emplID, legalEntity uuid.UUID) (..., error) {
// 	return s.stockwaveEngineService.GetStockOfLegalEntity(ctx, orgID, emplID, legalEntity)
// }
