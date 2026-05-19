package org_customer_service

import (
	org_customer_repo "sharpnr-engine/core/customer/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
)

type orgCustomerService struct {
	customerRepo org_customer_repo.OrgCustomerRepoInterface
	txManager    dbtx.TxManager
	seroEngine   *sero.SEROEngine
}

func NewOrgCustomerService(
	customerRepo org_customer_repo.OrgCustomerRepoInterface,
	txManager dbtx.TxManager,
	seroEngine *sero.SEROEngine,
) *orgCustomerService {
	return &orgCustomerService{
		customerRepo: customerRepo,
		txManager:    txManager,
		seroEngine:   seroEngine,
	}
}
