package sentinelAggregatorService

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Container chain resolution.
//
// A module answers only "what is my direct container". The aggregator walks the
// rest, because it is the only component that sees every registered type and
// therefore the only one that can follow HousedIn from a menu item up to its
// property.
//
// # Why the FULL chain, not just the direct parent
//
// Cedar's `in` is transitive, but only over entities that are actually present
// in the EntityMap. Verified against cedar-go v1.8.0:
//
//	room -> hotel -> property -> organization, every entity present
//	    `resource in Property::"p"`  => ALLOW  (two hops)
//
//	the same room, with the hotel entity omitted from the map
//	    `resource in Property::"p"`  => DENY
//
// The second case is silent: one-hop checks still pass, so a gap in the chain
// only shows up once someone scopes a policy to a container two levels up. That
// is why this returns every ancestor rather than the nearest one, and why the
// engine must stamp all of them.

// ResolveAncestry implements [SentinelAggregatorServiceInterface]: each
// instance's container chain, nearest first.
//
// The Organization root is deliberately NOT included. The engine stamps that
// edge itself from the request's org, so no module declares it and no walk has
// to reach it.
//
// Instances whose container cannot be resolved — deleted, or not visible to the
// caller — are simply absent from the returned map rather than erroring. A
// missing ancestry means container-scoped policies will not match, which is the
// safe direction to fail in.
func (s *sentinelAggregatorService) ResolveAncestry(
	ctx context.Context, orgID uuid.UUID, resourceType string, instanceIDs []string,
) (map[string][]ResourceRef, error) {
	if _, ok := s.resources[resourceType]; !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnknownResourceType, resourceType)
	}

	out := make(map[string][]ResourceRef, len(instanceIDs))
	if len(instanceIDs) == 0 {
		return out, nil
	}
	for _, id := range instanceIDs {
		out[id] = nil
	}

	// frontier maps a container instance id at the current level back to every
	// original instance descending from it, so one resolver call serves the
	// whole batch rather than one call per instance per level.
	frontier := make(map[string][]string, len(instanceIDs))
	for _, id := range instanceIDs {
		frontier[id] = []string{id}
	}

	currentType := resourceType
	for {
		entry := s.resources[currentType]
		if entry.housedIn == "" || entry.resolveContainer == nil {
			return out, nil // reached a top-level type; Organization is stamped by the engine
		}

		ids := make([]string, 0, len(frontier))
		for id := range frontier {
			ids = append(ids, id)
		}

		containers, err := entry.resolveContainer(ctx, orgID, ids)
		if err != nil {
			return nil, fmt.Errorf("resolving container of %q: %w", currentType, err)
		}
		if len(containers) == 0 {
			return out, nil
		}

		next := make(map[string][]string, len(containers))
		for childID, containerID := range containers {
			if containerID == "" {
				continue
			}
			ref := ResourceRef{Type: entry.housedIn, ID: containerID}
			for _, origin := range frontier[childID] {
				out[origin] = append(out[origin], ref)
			}
			next[containerID] = append(next[containerID], frontier[childID]...)
		}
		if len(next) == 0 {
			return out, nil
		}

		frontier = next
		currentType = entry.housedIn
		// No depth guard needed: validateHierarchy proved the containment graph
		// is acyclic at bootstrap, so this walk is bounded by the number of
		// registered types.
	}
}
