package kitchenservices

import (
	dbtx "sharpnr-engine/core/dbTx"
	kitchenrepo "sharpnr-engine/modules/hospitality/kitchen/repos"
	"sharpnr-engine/sero"
	sse_engine "sharpnr-engine/sse"
)

type kitchenService struct {
	kitchenRepo kitchenrepo.KitchenRepositoryInterface
	seroEngine  *sero.SEROEngine
	sseEngine   *sse_engine.StreamManager
	txManager   dbtx.TxManager
}

func NewKitchenService(
	kitchenRepo kitchenrepo.KitchenRepositoryInterface,
	seroEngine *sero.SEROEngine,
	sseEngine *sse_engine.StreamManager,
	txManager dbtx.TxManager,
) *kitchenService {
	return &kitchenService{
		kitchenRepo: kitchenRepo,
		seroEngine:  seroEngine,
		sseEngine:   sseEngine,
		txManager:   txManager,
	}
}
