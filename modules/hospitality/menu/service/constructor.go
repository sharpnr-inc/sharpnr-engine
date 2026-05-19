package menuservice

import (
	menurepo "sharpnr-engine/modules/hospitality/menu/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
	tax_engine_service "sharpnr-engine/service_engines/tax_engine/service"
)

type menuService struct {
	menuRepo  menurepo.MenuRepositoryInterface
	txManager dbtx.TxManager
	taxEngine tax_engine_service.TaxEngineServiceInterface
}

func NewHospitalityMenuService(
	menuRepo menurepo.MenuRepositoryInterface,
	txManager dbtx.TxManager,
	taxEngine tax_engine_service.TaxEngineServiceInterface,
) *menuService {
	return &menuService{
		menuRepo:  menuRepo,
		txManager: txManager,
		taxEngine: taxEngine,
	}

}
