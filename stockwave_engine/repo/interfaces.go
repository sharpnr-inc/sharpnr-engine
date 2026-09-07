package stockwaveEngineRepo

// StockwaveEngineRepoInterface is the persistence seam for the stockwave
// engine.
//
// It is deliberately empty: migration 000062 dropped the inventory tables this
// package used to read and write, and the account layer
// (stockwave_engine.org_stockwave_engine_accounts and
// stockwave_engine.stockwave_engine_account_settings) has no repository methods
// yet. The rewrite fills it in following the shape the sibling engines already
// use — see cortex_engine/repo/interface.go: initialize account, create its
// settings row, read by org, and the active/inactive switch.
type StockwaveEngineRepoInterface interface {
	
}
