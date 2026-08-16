// Package orchestrator_service implements the organization control panel's
// backend-for-frontend: a facade that composes the domain engines (properties,
// legal entities, payments, inventory, ...) into the flows the frontend needs.
//
// # Structure
//
// One small orchestrator per domain, each in its own file, each holding only
// the dependencies that domain needs (e.g. paymentsOrchestrator holds the
// payment engine and the tx manager, nothing else). The exported facade in
// constructor.go embeds all of them; method promotion makes the composed
// struct satisfy OrchestratorServiceInterface. To add a domain: create its
// file with a struct + constructor, add a sub-interface in interface.go, and
// embed it in orchestratorService.
//
// A domain that outgrows one file becomes a package instead — see
// orchestrator/service/property, which nests a sub-package per business domain
// under it. The facade embeds the package's exported orchestrator and nothing
// else changes; its sub-interface moves to interface_<domain>.go beside it.
//
// # Rules
//
//  1. Authorization: every orchestrator method calls guard.Authorize(...)
//     before doing work, UNLESS a comment names the domain-service check that
//     already covers it (see CreateProperty for the pattern). Silence is not
//     an exemption. The guard lives in orchestrator/service/authz so domain
//     sub-packages can share the one instance built in constructor.go.
//
//  2. Transactions: the orchestrator opens txManager.WithTx only when it
//     coordinates two or more writes across services. Single delegated calls
//     and reads pass ctx through untouched — the domain service owns its own
//     consistency.
//
//  3. Wire types: nothing crosses OrchestratorServiceInterface except
//     orchestrator_dtos types and primitives. Domain models and engine
//     contracts are mapped at this layer so engine refactors never become
//     breaking API changes. (Several older methods still violate this; new
//     code must not add violations.)
package orchestrator_service
