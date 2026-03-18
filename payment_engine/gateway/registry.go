package payment_gateway

import "fmt"

// -----------------------------------------------
// PROVIDER REGISTRY
// Registered once at bootstrap
// Resolved at runtime by provider code
// Adding a new gateway = register it here only
// -----------------------------------------------

type ProviderRegistry struct {
	providers map[string]GatewayProvider
}

func NewProviderRegistry(providers ...GatewayProvider) *ProviderRegistry {
	r := &ProviderRegistry{
		providers: make(map[string]GatewayProvider),
	}
	for _, p := range providers {
		r.providers[p.Code()] = p
	}
	return r
}

func (r *ProviderRegistry) Resolve(code string) (GatewayProvider, error) {
	p, ok := r.providers[code]
	if !ok {
		return nil, fmt.Errorf("payment gateway provider '%s' is not registered", code)
	}
	return p, nil
}

func (r *ProviderRegistry) IsRegistered(code string) bool {
	_, ok := r.providers[code]
	return ok
}

func (r *ProviderRegistry) RegisteredCodes() []string {
	codes := make([]string, 0, len(r.providers))
	for code := range r.providers {
		codes = append(codes, code)
	}
	return codes
}
