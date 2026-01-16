package sharpnr_policy_engine

// This is the ONLY thing other systems see
type ResolvedPolicy struct {
	AllowedModules map[string]bool            // bar, restaurant, hotel
	AllowedActions map[string]map[string]bool // bar -> orders:true
}
