package stockwaveEngineServices

// StockwaveEngineServicesInterface is the stockwave engine's public surface.
//
// It is deliberately empty: every method it used to carry read or wrote the
// inventory tables migration 000062 dropped. The rewrite fills it in starting
// with the account layer the sibling engines already expose — see
// cortex_engine/service/interface.go: start the service for an org, and the
// account-status read every other flow is gated on.
type StockwaveEngineServicesInterface interface {
}
