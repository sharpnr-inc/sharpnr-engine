package sharpCrmEngineService

import (
	auditEngineService "sharpnr-engine/audit_log_engine/service"
	org_customer_service "sharpnr-engine/core/customer/service"
	dbtx "sharpnr-engine/pkg/dbTx"
	sharpCrmEngineRepo "sharpnr-engine/sharpCRM_engine/repo"
)

type Service struct {

	// this will give us access to the organization customer records
	coreOrgCustomerService org_customer_service.OrgCustomerServiceInterface
	// then we will soon add the all Business Ops to this service to know what service they avail and purchased
	// possibly we will add the accounting service to know how much amount they spent and how much they owe

	crmRepo sharpCrmEngineRepo.SharpCrmEngineRepoInterace

	// txManager wraps a write and the audit entry that records it in one
	// transaction, so a rolled-back change never leaves a log line claiming it
	// happened. Provisioning needs it for a second reason: the account and its
	// settings row are one unit.
	txManager dbtx.TxManager

	// auditLogEngine records the account lifecycle onto the org's history. Nil
	// is a working configuration — recordAudit becomes a no-op — so a build
	// without the audit engine wired still serves the CRM API.
	auditLogEngine auditEngineService.AuditLogServiceInterface

	// mailerEngine
	// cables service (not now)
}

func NewSharpCrmEngineService(
	coreOrgCustomerService org_customer_service.OrgCustomerServiceInterface,
	crmRepo sharpCrmEngineRepo.SharpCrmEngineRepoInterace,
	txManager dbtx.TxManager,
	auditLogEngine auditEngineService.AuditLogServiceInterface,
) *Service {
	return &Service{
		coreOrgCustomerService: coreOrgCustomerService,
		crmRepo:                crmRepo,
		txManager:              txManager,
		auditLogEngine:         auditLogEngine,
	}
}
