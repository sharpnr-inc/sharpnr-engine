package menuservice

import (
	menurepo "sharpnr-engine/business_ops/hospitality/menu/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
)

type menuService struct {
	menuRepo  menurepo.MenuRepositoryInterface
	txManager dbtx.TxManager
}

func NewHospitalityMenuService(
	menuRepo menurepo.MenuRepositoryInterface,
	txManager dbtx.TxManager,

) *menuService {
	return &menuService{
		menuRepo:  menuRepo,
		txManager: txManager,
	}

}
