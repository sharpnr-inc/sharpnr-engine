package organization_services

import (
	organization_repo "sharpnr-engine/core/organization/repo"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sharp_mailer_engine"
)

type organizationService struct {
	mailerEngine *sharp_mailer_engine.SharpMailerEngine
	orgRedisRepo organization_repo.RedisOrganizationRepositoryInterface
	orgMainRepo  organization_repo.OrganizationRepositoryInterface
	txManager    dbtx.TxManager
}

func NewOrganizationService(
	mailerEngine *sharp_mailer_engine.SharpMailerEngine,
	orgRedisRepo organization_repo.RedisOrganizationRepositoryInterface,
	orgMainRepo organization_repo.OrganizationRepositoryInterface,
	txManager dbtx.TxManager,

) *organizationService {
	return &organizationService{
		mailerEngine: mailerEngine,
		orgRedisRepo: orgRedisRepo,
		orgMainRepo:  orgMainRepo,
		txManager:    txManager,
	}
}
