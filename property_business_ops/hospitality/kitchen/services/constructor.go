package kitchenservices

import (
	kitchenrepo "sharpnr-engine/property_business_ops/hospitality/kitchen/repos"
	dbtx "sharpnr-engine/pkg/dbTx"
	"sharpnr-engine/sero"
	sse_engine "sharpnr-engine/sse_engine"
	"sync"
)

type kitchenService struct {
	kitchenRepo kitchenrepo.KitchenRepositoryInterface
	seroEngine  *sero.SEROEngine
	sseEngine   *sse_engine.StreamManager
	txManager   dbtx.TxManager

	// Per-kitchen locks (uuid.UUID -> *sync.Mutex) serializing SSE
	// broadcasts; see broadcastKots.
	kotBroadcastLocks sync.Map
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
