package app_test

import (
	"sharpnr-engine/app"
	dbtx "sharpnr-engine/core/dbTx"
	"sharpnr-engine/sere"
	"testing"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func TestBootstrapModules(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		db         *gorm.DB
		redis      *redis.Client
		txManager  dbtx.TxManager
		sereEngine *sere.SERE
		want       *app.AppModules
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := app.BootstrapModules(tt.db, tt.redis, tt.txManager, tt.sereEngine)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("BootstrapModules() = %v, want %v", got, tt.want)
			}
		})
	}
}
