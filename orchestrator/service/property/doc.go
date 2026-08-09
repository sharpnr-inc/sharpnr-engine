// Package property_orchestrator owns the /property route tree: the property
// itself, plus the composition of the business modules, infra modules and
// resources attached to it.
//
// # Structure
//
//	orchestrator.go       PropertyOrchestrator + New — the wiring
//	base.go               the property itself (list / details / create)
//	business_modules.go   attach + list business operation modules
//	infra_modules.go      attach + list infra modules
//	hospitality/          the hospitality domain's flows (sub-package)
//
// The property files stay composition-only: authorize, ask each business
// domain for its slice, assemble the response DTO. Everything that knows how a
// domain actually works lives in that domain's sub-package, which holds only
// that domain's services. A second domain (healthcare, retail) is a sibling of
// hospitality/ plus one embedded field on PropertyOrchestrator — the files here
// gain a few lines, not a new concern.
//
// # Promotion
//
// PropertyOrchestrator is embedded in orchestrator_service.orchestratorService,
// and embeds its domain orchestrators in turn, so a method defined in
// hospitality/ is promoted all the way to the composed facade with no wiring.
// The cost of that convenience is a single flat method namespace across every
// orchestrator: keep method names entity-qualified (GetHotelDetails, not
// GetDetails) or the promoted selector becomes ambiguous.
//
// The rules in orchestrator/service/doc.go apply unchanged: guard first,
// txManager only when coordinating multiple writes, orchestrator_dtos on the
// wire.
package property_orchestrator
