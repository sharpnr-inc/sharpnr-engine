package menuservice

import (
	dbtx "sharpnr-engine/core/dbTx"
	menurepo "sharpnr-engine/modules/hospitality/menu/repo"
)

type menuService struct {
	menuRepo  menurepo.MenuRepositoryInterface
	txManager dbtx.TxManager
}

func NewHospitalityMenuService(menuRepo menurepo.MenuRepositoryInterface, txManager dbtx.TxManager) *menuService {
	return &menuService{
		menuRepo:  menuRepo,
		txManager: txManager,
	}

}
