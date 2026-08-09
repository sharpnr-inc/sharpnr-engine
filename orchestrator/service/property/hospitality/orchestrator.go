// Package hospitality_orchestrator owns every property-scoped hospitality flow:
// the hospitality slice of a property's business/infra module composition, plus
// the per-entity operations (hotel, restaurant, bar, kitchen, menu) mounted
// under /property/:propertyID/... in the route tree.
//
// It is a sub-package of property_orchestrator rather than a file inside it so
// the hospitality flows can only reach the hospitality domain — orgPropertyService
// and organizationService are deliberately out of scope here. A second business
// domain (healthcare, retail) gets a sibling package and one embedded field in
// property_orchestrator.PropertyOrchestrator.
//
// Rules from orchestrator/service/doc.go apply unchanged: guard first, tx only
// when coordinating multiple writes, orchestrator_dtos on the wire.
package hospitality_orchestrator

import (
	hospitalitydomain "sharpnr-engine/app/hospitality_domain"
	orchestrator_authz "sharpnr-engine/orchestrator/service/authz"
)

// HospitalityOrchestrator implements the hospitality portion of the property
// orchestration interface. Its methods are promoted through
// property_orchestrator.PropertyOrchestrator up to the composed orchestrator
// facade, so a new per-entity method needs no wiring beyond being defined here.
//
// Method names must stay entity-qualified (GetHotelDetails, not GetDetails):
// they share a flat namespace with every other orchestrator on the facade, and
// a collision makes the promoted selector ambiguous.
type HospitalityOrchestrator struct {
	guard *orchestrator_authz.Guard

	// domainSvc is the hospitality domain's service registry (hotel, bar,
	// restaurant, kitchen, menu, order engine).
	domainSvc *hospitalitydomain.HospitalityDomain
}

func New(
	guard *orchestrator_authz.Guard,
	domainSvc *hospitalitydomain.HospitalityDomain,
) *HospitalityOrchestrator {
	return &HospitalityOrchestrator{
		guard:     guard,
		domainSvc: domainSvc,
	}
}
