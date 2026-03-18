// ProviderRegistry holds all registered gateway providers.
// Built once at application startup in bootstrapPaymentEngine().
// Resolved at runtime by provider code string.
//
// USAGE:
//   // Bootstrap (app/bootstrap_payment_engine.go)
//   registry := payment_gateway.NewProviderRegistry(
//       cash.NewCashProvider(),
//       phonepe.NewPhonePeProvider(),
//   )
//
//   // Runtime (gateway_resolver.go)
//   provider, err := registry.Resolve("PHONEPE")
//
// EXTENSIBILITY:
//   Adding Razorpay: create gateway/razorpay/provider.go, implement GatewayProvider,
//   register razorpay.NewRazorpayProvider() here. Zero changes anywhere else.

package payment_gateway

import "fmt"

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

// Resolve returns the provider for the given code.
// Returns error if provider is not registered.
// Called by gateway_resolver.go — never called directly by payment service.
func (r *ProviderRegistry) Resolve(code string) (GatewayProvider, error) {
	p, ok := r.providers[code]
	if !ok {
		return nil, fmt.Errorf("payment gateway provider '%s' is not registered", code)
	}
	return p, nil
}

// IsRegistered checks if a provider code is available.
// Used by gateway config service to validate before saving tenant config.
func (r *ProviderRegistry) IsRegistered(code string) bool {
	_, ok := r.providers[code]
	return ok
}

// RegisteredCodes returns all registered provider codes.
// Used by the UI to show available payment methods to configure.
func (r *ProviderRegistry) RegisteredCodes() []string {
	codes := make([]string, 0, len(r.providers))
	for code := range r.providers {
		codes = append(codes, code)
	}
	return codes
}
